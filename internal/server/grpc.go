package server

import (
	v1 "YYeTsBot-Go/api/helloworld/v1"
	ets "YYeTsBot-Go/api/yyets/v1"
	"YYeTsBot-Go/internal/conf"
	"YYeTsBot-Go/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, resource *service.ResourceService, comment *service.CommentService, announcement *service.AnnouncementService, metricsService *service.MetricsService, adsense *service.AdsenseService, douban *service.DoubanService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
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
	v1.RegisterGreeterServer(srv, greeter)
	ets.RegisterResourceServer(srv, resource)
	ets.RegisterCommentServer(srv, comment)
	ets.RegisterAnnouncementServer(srv, announcement)
	ets.RegisterMetricsServer(srv, metricsService)
	ets.RegisterAdsenseServer(srv, adsense)
	ets.RegisterDoubanServer(srv, douban)
	return srv
}
