// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"tinytiktok/cmd/api/biz/rpc"
)

func Init() {
	binding.SetLooseZeroMode(true)
	rpc.InitRPC()
	// hlog init
	hlog.SetLogger(logrus.NewLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
		server.WithMaxRequestBodySize(500*1024*1024),
		tracer,
	)
	// use pprof mw
	pprof.Register(h)
	// use otel mw
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
