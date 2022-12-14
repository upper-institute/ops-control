package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	sdtypes "github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	cachetypes "github.com/envoyproxy/go-control-plane/pkg/cache/types"
	parameterapi "github.com/upper-institute/ops-control/gen/api/parameter"
	service_discovery "github.com/upper-institute/ops-control/gen/api/service-discovery"
	domainregistry "github.com/upper-institute/ops-control/internal/domain-registry"
	parameter "github.com/upper-institute/ops-control/internal/parameter"
	sdinternal "github.com/upper-institute/ops-control/internal/service-discovery"
	"go.uber.org/zap"
)

type cloudMapServiceTag_ParameterPathProvider struct {
	v string
}

func (c *cloudMapServiceTag_ParameterPathProvider) FromTags(tagName string, tags []sdtypes.Tag) {
	for _, tag := range tags {
		if aws.ToString(tag.Key) == tagName {
			c.v = aws.ToString(tag.Value)
			break
		}
	}
}

func (c *cloudMapServiceTag_ParameterPathProvider) GetParameterPath() string {
	return c.v
}

type CloudMapServiceDiscovery struct {
	namespacesNames  []string
	parameterPathTag string
	xdsClusterName   string

	cloudMapClient *servicediscovery.Client

	logger *zap.SugaredLogger

	parameterStore          parameter.ParameterStore
	parameterFileDownloader parameter.ParameterFileDownloader
	domainRegistry          domainregistry.DomainRegistryService
}

func NewCloudMapServiceDiscovery(
	namespacesNames []string,
	parameterPathTag string,
	xdsClusterName string,
	cloudMapClient *servicediscovery.Client,
	logger *zap.SugaredLogger,
) sdinternal.ServiceDiscoveryService {

	return &CloudMapServiceDiscovery{
		namespacesNames:  namespacesNames,
		parameterPathTag: parameterPathTag,
		xdsClusterName:   xdsClusterName,
		cloudMapClient:   cloudMapClient,
		logger:           logger.With("xds_cluster", xdsClusterName, "namespaces_names", namespacesNames),
	}

}

func (c *CloudMapServiceDiscovery) getListNamespacesInputFilters() []sdtypes.NamespaceFilter {

	c.logger.Debugw("Building NamespaceFilter to get namespaces IDs from namespaces names")

	namesFilter := sdtypes.NamespaceFilter{
		Name:      sdtypes.NamespaceFilterNameName,
		Values:    c.namespacesNames,
		Condition: sdtypes.FilterConditionEq,
	}

	return []sdtypes.NamespaceFilter{namesFilter}

}

func (c *CloudMapServiceDiscovery) getListServicesInputFilters(ctx context.Context) ([]sdtypes.ServiceFilter, error) {

	listNamespacesReq := servicediscovery.NewListNamespacesPaginator(
		c.cloudMapClient,
		&servicediscovery.ListNamespacesInput{
			Filters: c.getListNamespacesInputFilters(),
		},
	)

	namespaceIdsFilter := sdtypes.ServiceFilter{
		Name:      sdtypes.ServiceFilterNameNamespaceId,
		Values:    []string{},
		Condition: sdtypes.FilterConditionEq,
	}

	for listNamespacesReq.HasMorePages() {

		listNamespacesPage, err := listNamespacesReq.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, namespace := range listNamespacesPage.Namespaces {

			namespaceId := aws.ToString(namespace.Id)

			c.logger.Infow("Found namespace from AWS Cloud Map", "namespace_id", namespaceId)

			namespaceIdsFilter.Values = append(namespaceIdsFilter.Values, namespaceId)

		}

	}

	return []sdtypes.ServiceFilter{namespaceIdsFilter}, nil

}

func (c *CloudMapServiceDiscovery) setupIngress(ctx context.Context, sdState *sdinternal.ServiceDiscoveryState, paramSet *parameter.ParameterSet) error {

	ingressConfig := &service_discovery.Ingress{}

	err := paramSet.ParseProtoJson(ctx, parameterapi.WellKnown_WELL_KNOWN_INGRESS.String(), ingressConfig)
	if err != nil {
		return err
	}

	ingressConfig.XdsClusterName = c.xdsClusterName

	switch {

	case len(ingressConfig.Domains) == 0:
		c.logger.Warnw("")

	case c.domainRegistry == nil:
		c.logger.Errorw("")

	default:

		for i, domain := range ingressConfig.Domains {

			c.logger.Debugw("")

			err := c.domainRegistry.RegisterIngressDomain(ctx, domain)
			if err != nil {
				c.logger.Errorw(err.Error(), "domain_index", i)
				return err
			}

			c.logger.Infow("")

		}
	}

	sdState.AddIngress(ingressConfig)

	return nil

}

func (c *CloudMapServiceDiscovery) setupServiceCluster(ctx context.Context, sdState *sdinternal.ServiceDiscoveryState, paramSet *parameter.ParameterSet, service sdtypes.ServiceSummary) error {

	serviceName := aws.ToString(service.Name)

	addServiceCluster := &service_discovery.ServiceCluster{}

	err := paramSet.ParseProtoJson(ctx, parameterapi.WellKnown_WELL_KNOWN_SERVICE_CLUSTER.String(), addServiceCluster)
	if err != nil {
		return err
	}

	addServiceCluster.ServiceClusterName = serviceName
	addServiceCluster.XdsClusterName = c.xdsClusterName

	sdState.AddServiceCluster(addServiceCluster)

	listInstancesReq := servicediscovery.NewListInstancesPaginator(
		c.cloudMapClient,
		&servicediscovery.ListInstancesInput{
			ServiceId: service.Id,
		},
	)

	addServiceEndpointsInput := &service_discovery.ServiceEndpoints{
		ServiceClusterName: serviceName,
		Endpoints:          []*service_discovery.Endpoint{},
	}

	for listInstancesReq.HasMorePages() {

		listInstancesPage, err := listInstancesReq.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, instance := range listInstancesPage.Instances {

			address, ok := instance.Attributes["AWS_INSTANCE_IPV4"]
			if !ok {
				c.logger.Infow("Instance without AWS_INSTANCE_IPV4 key", "service_name", serviceName)
				continue
			}

			c.logger.Infow("Add endpoint from AWS Cloud Map service instance", "service_name", serviceName, "instance_ipv4", address, "upstream_port", addServiceCluster.UpstreamPort)

			addServiceEndpointsInput.Endpoints = append(
				addServiceEndpointsInput.Endpoints,
				&service_discovery.Endpoint{
					Protocol:  corev3.SocketAddress_TCP,
					Address:   address,
					PortValue: addServiceCluster.UpstreamPort,
				},
			)

		}

	}

	sdState.AddServiceEndpoints(addServiceEndpointsInput)

	return nil

}

func (c *CloudMapServiceDiscovery) discoverService(ctx context.Context, sdState *sdinternal.ServiceDiscoveryState, service sdtypes.ServiceSummary) error {

	paramSet := parameter.NewParameterSet(c.parameterFileDownloader, c.logger)

	serviceName := aws.ToString(service.Name)

	c.logger.Debugw("Starting service discovery process (AWS Cloud Map)", "service_name", serviceName)

	listServiceTagsRes, err := c.cloudMapClient.ListTagsForResource(ctx, &servicediscovery.ListTagsForResourceInput{ResourceARN: service.Arn})
	if err != nil {
		return err
	}

	paramPathProvider := &cloudMapServiceTag_ParameterPathProvider{}
	paramPathProvider.FromTags(c.parameterPathTag, listServiceTagsRes.Tags)

	c.logger.Debugw("Load parameter path from service tag", "service_name", serviceName, "tag_key", c.parameterPathTag, "parameter_path_value", paramPathProvider.GetParameterPath())

	if len(paramPathProvider.GetParameterPath()) == 0 {
		c.logger.Infow("Ignoring service discovery 'case parameter_path is empty", "service_name", serviceName)
		return nil
	}

	err = c.parameterStore.Load(ctx, paramPathProvider, paramSet)
	if err != nil {
		return err
	}

	wnIngress := parameterapi.WellKnown_WELL_KNOWN_INGRESS.String()
	wnServiceCluster := parameterapi.WellKnown_WELL_KNOWN_SERVICE_CLUSTER.String()

	switch {

	case paramSet.HasFile(wnIngress):
		c.logger.Infow("Ingress configuration file found", "service_name", serviceName)
		return c.setupIngress(ctx, sdState, paramSet)

	case paramSet.HasFile(wnServiceCluster):
		c.logger.Infow("Service cluster configuration file found", "service_name", serviceName)
		return c.setupServiceCluster(ctx, sdState, paramSet, service)

	}

	c.logger.Infow("No configuration file found", "service_name", serviceName)

	return nil

}

func (c *CloudMapServiceDiscovery) SetParameterStore(parameterStore parameter.ParameterStore) {
	c.parameterStore = parameterStore
}

func (c *CloudMapServiceDiscovery) SetParameterFileDownloader(parameterFileDownloader parameter.ParameterFileDownloader) {
	c.parameterFileDownloader = parameterFileDownloader
}

func (c *CloudMapServiceDiscovery) SetDomainRegistry(domainRegistry domainregistry.DomainRegistryService) {
	c.domainRegistry = domainRegistry
}

func (c *CloudMapServiceDiscovery) checkDependencies() error {

	if c.parameterFileDownloader == nil || c.parameterStore == nil {
		return fmt.Errorf("AWS Cloud Map service discovery requires a parameter store and a parameter file downloader")
	}

	return nil

}

func (c *CloudMapServiceDiscovery) Discover(ctx context.Context) (map[string][]cachetypes.Resource, error) {

	err := c.checkDependencies()
	if err != nil {
		return nil, err
	}

	listServicesFilters, err := c.getListServicesInputFilters(ctx)
	if err != nil {
		return nil, err
	}

	listServicesReq := servicediscovery.NewListServicesPaginator(
		c.cloudMapClient,
		&servicediscovery.ListServicesInput{
			Filters: listServicesFilters,
		},
	)

	sdState := sdinternal.NewServiceDiscoveryState(c.logger)

	for listServicesReq.HasMorePages() {

		listServicesPage, err := listServicesReq.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, service := range listServicesPage.Services {

			err = c.discoverService(ctx, sdState, service)
			if err != nil {
				return nil, err
			}

		}

	}

	c.logger.Infow("Build envoy service discovery resources map")

	err = sdState.Build()
	if err != nil {
		return nil, err
	}

	return sdState.Resources, nil

}
