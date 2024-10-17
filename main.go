package main

import (
	"github.com/fnless/otel-examples/cmd"
	_ "github.com/fnless/otel-examples/pkg/otel"
)

func main() {
	cmd.Execute()
}
