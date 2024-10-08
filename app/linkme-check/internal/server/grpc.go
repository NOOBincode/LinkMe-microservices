package server

import (
	v1 "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-check/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, checkSvc *service.CheckService, logger log.Logger, tp *tracesdk.TracerProvider) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			// 追踪中间件：用于分布式追踪，集成 OpenTelemetry
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Client(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterCheckServer(srv, checkSvc)
	return srv
}
