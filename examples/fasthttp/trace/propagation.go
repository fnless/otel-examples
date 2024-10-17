package trace

import (
	"context"

	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
)

func Extract(header *fasthttp.RequestHeader) context.Context {
	propagator := otel.GetTextMapPropagator()
	carrier := HeaderCarrier{header}
	return propagator.Extract(context.Background(), carrier)
}

func Inject(ctx context.Context, header *fasthttp.RequestHeader) {
	propagator := otel.GetTextMapPropagator()
	carrier := HeaderCarrier{header}
	propagator.Inject(ctx, carrier)
}
