package envoy

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/upper-institute/ops-control/cmd/controllers/parameter"
	"github.com/upper-institute/ops-control/internal/logger"
	sdinternal "github.com/upper-institute/ops-control/internal/service-discovery"
	"github.com/upper-institute/ops-control/providers/aws"
	"github.com/upper-institute/ops-control/providers/envoy"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
)

var (
	xdsServer serverv3.Server

	xdsServerCmd = &cobra.Command{
		Use:   "xds-server",
		Short: "Run xDS server",
		RunE: func(cmd *cobra.Command, args []string) error {

			var (
				envoyConfiguration = envoy.NewConfiguration(logger.SugaredLogger)
				cache              = cache.NewSnapshotCache(false, cache.IDHash{}, nil)

				serviceDiscoveryService sdinternal.ServiceDiscoveryService

				discoveryMinInterval = viper.GetDuration("envoy.discoveryMinInterval")
				nodeId               = viper.GetString("envoy.nodeId")
				xdsClusterName       = viper.GetString("envoy.xdsCluster.name")
				parameterPathTag     = viper.GetString("envoy.parameter.pathTag")

				log = logger.SugaredLogger
			)

			ctx := context.Background()

			parameterStore, parameterFileDownloader, err := parameter.LoadParameterProviders(ctx)
			if err != nil {
				return err
			}

			switch {
			case viper.GetBool("envoy.aws.cloudMap"):

				log.Infow("Envoy AWS Cloud Map xDS Server enabled")

				config, err := config.LoadDefaultConfig(ctx)
				if err != nil {
					return err
				}

				cloudMapClient := servicediscovery.NewFromConfig(config)

				serviceDiscoveryService = aws.NewCloudMapServiceDiscovery(
					viper.GetStringSlice("envoy.aws.cloudMap.namespaces"),
					parameterPathTag,
					xdsClusterName,
					cloudMapClient,
					logger.SugaredLogger,
					parameterStore,
					parameterFileDownloader,
				)

			}

			go func() {

				for {

					ctx := context.Background()

					log.Info("Starting discovery cycle")

					resources, err := serviceDiscoveryService.Discover(ctx)
					if err != nil {
						log.Fatalw(err.Error())
					}

					log.Info("End of discovery cycle")

					envoyConfiguration.Resources = resources

					snapshot, err := envoyConfiguration.DoSnapshot()
					if err != nil {
						log.Fatalw(err.Error())
					}

					err = cache.SetSnapshot(ctx, nodeId, snapshot)
					if err != nil {
						log.Fatalw(err.Error())
					}

					log.Infow("Discovery process interval", "interval_duration", discoveryMinInterval)

					time.Sleep(discoveryMinInterval)
				}

			}()

			xdsServer = serverv3.NewServer(ctx, cache, nil)

			return nil
		},
	}
)
