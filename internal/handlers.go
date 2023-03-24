package internal

import (
	"encoding/json"
	"net/http"
	"time"
	"tokay/configs"
)

func generateTokens(res http.ResponseWriter, req *http.Request) {
	subject := req.URL.Query().Get("sub")
	if subject == "" {
		http.Error(res, "Missing sub as query parameter", http.StatusBadRequest)
		return
	}

	lifeTime, _ := time.ParseDuration(configs.Env.ACCESS_TOKEN_LIFETIME)
	accessToken, err := SignToken(subject, "tokay", configs.Env.ACCESS_TOKEN_SECRET, lifeTime)
	if err != nil {
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	lifeTime, _ = time.ParseDuration(configs.Env.REFRESH_TOKEN_LIFETIME)
	refreshToken, err := SignToken(subject, "tokay", configs.Env.REFRESH_TOKEN_SECRET, lifeTime)
	if err != nil {
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func blacklistToken(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	subject, exists := ctx.Value("subject").(string)
	if !exists {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	token, exists := ctx.Value("token").(string)
	if !exists {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	if err := store.Blacklist(subject, token); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("ok"))
}

func checkBlacklist(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	subject, exists := ctx.Value("subject").(string)
	if !exists {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	token, exists := ctx.Value("token").(string)
	if !exists {
		http.Error(res, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	is, err := store.IsBlacklisted(subject, token)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if is {
		http.Error(res, "Token is in the blacklist", http.StatusUnauthorized)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("ok"))
}
