package app

import (
	"context"
	"net/http"
	"strings"
	"tokay/configs"
)

type key string

type Verifier struct {
	secret string
}

func NewVerifier(secret string) *Verifier {
	return &Verifier{secret: secret}
}

func (verifier *Verifier) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")
		if auth == "" {
			http.Error(res, "Autorization header is not provided", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			http.Error(res, "Could not find Bearer token in Authorization header", http.StatusForbidden)
			return
		}

		claims, err := DecodeToken(token, verifier.secret)
		if err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}

		parentCtx := req.Context()
		ctx := context.WithValue(parentCtx, key("token"), token)
		ctx = context.WithValue(ctx, key("subject"), claims["sub"].(string))
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

var accessTokenVerifier = NewVerifier(configs.Env.ACCESS_TOKEN_SECRET)
var refreshTokenVerifier = NewVerifier(configs.Env.ACCESS_TOKEN_SECRET)
