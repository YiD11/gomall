package serversuite

import (
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	// prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	// "github.com/YiD11/gomall/common/mtl"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       []string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// server.WithTracer(prometheus.NewServerTracer(
		// 	"",
		// 	"",
		// 	prometheus.WithDisableServer(true),
		// 	prometheus.WithRegistry(mtl.Registry),
		// )),

		server.WithSuite(tracing.NewServerSuite()),
	}

	// registry
	r, err := consul.NewConsulRegister(s.RegistryAddr[0])
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
