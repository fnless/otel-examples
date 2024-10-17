package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var (
	Endpoint    = "localhost:4318"
	ServiceName = "otel-examples"
)

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// Set up the *otlphttp* exporter
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(Endpoint),
		otlptracehttp.WithInsecure(),
	)
	return otlptrace.New(ctx, client)
}

func newResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(ServiceName),
		),
	)
}

func newTraceProvider(exp sdktrace.SpanExporter, res *resource.Resource) *sdktrace.TracerProvider {
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp,
			sdktrace.WithMaxExportBatchSize(8192),
			sdktrace.WithMaxQueueSize(65536)),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
	)
	return tracerProvider
}

func init() {
	ctx := context.Background()
	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}
	res, err := newResource(ctx)
	if err != nil {
		log.Fatal(err)
	}
	tp := newTraceProvider(exp, res)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}))
}
