package mtl

import (
	"log"
	"os"
	"strconv"

	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

var (
	p provider.OtelProvider
)

func InitTracing(serviceName string) provider.OtelProvider {

	insecure, err := strconv.ParseBool(os.Getenv("OTEL_EXPORTER_OTLP_INSECURE"))
	if err != nil {
		log.Panicln("Wrong type of environment variable OTEL_EXPORTER_OTLP_INSECURE:", err)
	}

	options := []provider.Option{
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		provider.WithEnableMetrics(false),
	}
	if insecure {
		options = append(options, provider.WithInsecure())
	}
	p = provider.NewOpenTelemetryProvider(options...)
	return p
}
