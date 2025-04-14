package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

type contextKey string

const ClientIPKey contextKey = "client-ip"

func ClientIPMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ip := extractClientIP(ctx)
			ctx = context.WithValue(ctx, ClientIPKey, ip)
			return handler(ctx, req)
		}
	}
}
func extractClientIP(ctx context.Context) string {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return ""
	}

	switch t := tr.(type) {
	case *http.Transport:
		req := t.Request()
		// 优先从 X-Forwarded-For
		if xff := req.Header.Get("X-Forwarded-For"); xff != "" {
			return strings.TrimSpace(strings.Split(xff, ",")[0])
		}
		if xrip := req.Header.Get("X-Real-IP"); xrip != "" {
			return xrip
		}
		host, _, err := net.SplitHostPort(req.RemoteAddr)
		if err == nil {
			return host
		}
		return req.RemoteAddr

	case *grpc.Transport:
		if p, ok := peer.FromContext(ctx); ok {
			host, _, err := net.SplitHostPort(p.Addr.String())
			if err != nil {
				return p.Addr.String()
			}
			return host
		}
	}
	return ""
}
