package mtl

import (
	"log"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName string, metricsPort string, registryAddr []string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	r, err := consul.NewConsulRegister(registryAddr[0])
	if err != nil {
		log.Fatalln(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)

	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr: addr,
		Weight: 1,
		Tags: map[string]string{"service": serviceName},
	}

	err = r.Register(registryInfo)
	if err != nil {
		log.Panicln(err)
	}
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	go http.ListenAndServe(metricsPort, nil)

	return r, registryInfo
}
