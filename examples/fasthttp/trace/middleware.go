package trace

import (
	"github.com/valyala/fasthttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// Trace is a fasthttp request middleware to export tracing data.
func Trace(handler fasthttp.RequestHandler, spanName string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// extract traceparent from request, and then use it to create a root span.
		rootSpanCtx, rootSpan := tracer.Start(
			Extract(&ctx.Request.Header),
			spanName,
			oteltrace.WithSpanKind(oteltrace.SpanKindServer), // tag: span.kind = server.
		)

		// inject propagation context into request
		Inject(rootSpanCtx, &ctx.Request.Header)
		defer rootSpan.End()
		handler(ctx)
		rootSpan.SetAttributes(semconv.HTTPResponseStatusCode(ctx.Response.StatusCode()))
	}
}
