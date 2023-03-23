package main

import (
	"net/http"
	"tokay/app"
	"tokay/configs"
)

func main() {
	if err := configs.InitConfigs(); err != nil {
		panic(err)
	}
	http.ListenAndServe(configs.Env.PORT, app.Router())
}
