package fasthttp

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	fasthttptrace "github.com/fnless/otel-examples/examples/fasthttp/trace"
	"github.com/fnless/otel-examples/pkg/service"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Propagation(ctx *fasthttp.RequestCtx) {
	hc := &fasthttp.HostClient{
		Addr: "localhost:8080", // The host address and port must be set explicitly
	}
	url := fasthttp.AcquireURI()
	url.Parse(nil, []byte("http://localhost:8080/"))
	req := fasthttp.AcquireRequest()
	req.SetURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	// inject propagation context
	ctxParent := fasthttptrace.Extract(&ctx.Request.Header)
	fasthttptrace.Inject(ctxParent, &req.Header)

	resp := fasthttp.AcquireResponse()
	err := hc.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		ctx.WriteString(err.Error())
		ctx.SetStatusCode(500)
		return
	}
	ctx.SetBody(resp.Body())
	ctx.SetStatusCode(resp.StatusCode())
	fasthttp.ReleaseResponse(resp)
}

type server struct{}

var _ service.Service = &server{}

func (s *server) Start() {
	r := router.New()
	r.GET("/", fasthttptrace.Trace(Index, "/index"))
	r.GET("/propagation", fasthttptrace.Trace(Propagation, "/propagation"))
	fmt.Println("examples/fasthttp server started at :8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}

func init() {
	service.Register("fasthttp", &server{})
}
