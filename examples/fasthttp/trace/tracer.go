package trace

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("otel-examples/fasthttp")
)
