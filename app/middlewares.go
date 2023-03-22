package app

import (
	"net/http"
)

func VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	})
}

func VerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	})
}
