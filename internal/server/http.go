package server

import (
	v1 "YYeTsBot-Go/api/helloworld/v1"
	ets "YYeTsBot-Go/api/yyets/v1"
	"YYeTsBot-Go/internal/conf"
	"YYeTsBot-Go/internal/middleware"
	"YYeTsBot-Go/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.yyets.v1.User/GetUser"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return true
		}
		return false
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, resource *service.ResourceService, comment *service.CommentService, announcement *service.AnnouncementService, metricsService *service.MetricsService, adsense *service.AdsenseService, douban *service.DoubanService, captchaService *service.CaptchaService, userService *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.ClientIPMiddleware(),
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte("your-secret-key"), nil
				}),
			).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	ets.RegisterResourceHTTPServer(srv, resource)
	ets.RegisterCommentHTTPServer(srv, comment)
	ets.RegisterAnnouncementHTTPServer(srv, announcement)
	ets.RegisterMetricsHTTPServer(srv, metricsService)
	ets.RegisterAdsenseHTTPServer(srv, adsense)
	ets.RegisterDoubanHTTPServer(srv, douban)
	ets.RegisterCaptchaHTTPServer(srv, captchaService)
	ets.RegisterUserHTTPServer(srv, userService)
	return srv
}
