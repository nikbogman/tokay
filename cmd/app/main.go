package main

import (
	"net/http"
	"tokay/configs"
	app "tokay/internal"
)

func main() {
	if err := configs.InitConfigs(); err != nil {
		panic(err)
	}
	http.ListenAndServe(configs.Env.PORT, app.Router())
}
