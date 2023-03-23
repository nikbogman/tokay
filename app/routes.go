package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/generate", GenerateTokens)
	router.Route("/verify", func(route chi.Router) {
		route.With(accessTokenVerifier.VerifyToken).Get("/access", CheckBlacklist)
		route.With(refreshTokenVerifier.VerifyToken).Get("/refresh",
			func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(http.StatusOK)
				res.Write([]byte("ok"))
			})
	})
	router.With(refreshTokenVerifier.VerifyToken).Get("/rotate", GenerateTokens)
	router.With(accessTokenVerifier.VerifyToken).Delete("/blacklist", BlacklistToken)
	return router
}
