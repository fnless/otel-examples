package cmd

import (
	"github.com/spf13/cobra"

	_ "github.com/fnless/otel-examples/examples/fasthttp"
	_ "github.com/fnless/otel-examples/examples/grpc"
	"github.com/fnless/otel-examples/pkg/service"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an example",
	Run: func(cmd *cobra.Command, args []string) {
		svc := "fasthttp"
		if len(args) > 0 {
			svc = args[0]
		}
		service.New(svc).Start()
	},
}
