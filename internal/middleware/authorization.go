package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.yyets.v1.User/GetUser"] = struct{}{}
	whiteList["/api.yyets.v1.Comment/CreateComment"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return true
		}
		return false
	}
}

func AuthorizationMiddleware(secretKey string) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return []byte(secretKey), nil
		}),
	).Match(NewWhiteListMatcher()).Build()
}
