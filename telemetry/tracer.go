package telemetry

import (
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/sdk/export/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.uber.org/zap"
)

// InitTracer creates an opentelemetry tracer pipepline
func InitTracer() {
	var exporter trace.SpanExporter
	var err error
	// TODO set another type of target exporter (jaeger or datadog, ...) for non local env
	exporter, err = stdout.NewExporter(stdout.WithPrettyPrint())
	if err != nil {
		zap.L().Panic("Failed to create stdout trace exporter", zap.Error(err))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithConfig(sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		}),
		sdktrace.WithSyncer(exporter),
	)
	otel.SetTracerProvider(tp)
}
