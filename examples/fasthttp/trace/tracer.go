package trace

import (
	"go.opentelemetry.io/otel"
)

var (
	// 虽然这里的 tracer 会在 pkg/otel.init() 初始化之前
	// 但是在 pkg/otel.init() 之后，tracer 会被 DelegateTracerProvider 代理
	// 详情参见:
	// - https://github.com/open-telemetry/opentelemetry-go/blob/v1.35.0/internal/global/state.go#L109-L112
	// - https://github.com/open-telemetry/opentelemetry-go/blob/v1.35.0/internal/global/trace.go#L68-L72
	// - https://github.com/open-telemetry/opentelemetry-go/blob/v1.35.0/internal/global/trace.go#L137-L139
	// - https://github.com/open-telemetry/opentelemetry-go/blob/v1.35.0/internal/global/trace.go#L144-L147
	tracer = otel.Tracer("otel-examples/fasthttp")
)
