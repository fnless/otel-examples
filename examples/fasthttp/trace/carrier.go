package trace

import (
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel/propagation"
)

type HeaderCarrier struct {
	header *fasthttp.RequestHeader
}

// Compile time check that MapCarrier implements the TextMapCarrier.
var _ propagation.TextMapCarrier = HeaderCarrier{}

// Get returns the value associated with the passed key.
func (c HeaderCarrier) Get(key string) string {
	return string(c.header.Peek(key))
}

// Set stores the key-value pair.
func (c HeaderCarrier) Set(key, value string) {
	c.header.Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (c HeaderCarrier) Keys() []string {
	var keys []string
	c.header.VisitAll(func(key, _ []byte) {
		keys = append(keys, string(key))
	})
	return keys
}
