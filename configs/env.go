package configs

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var Env envVars

func InitConfigs() (err error) {
	Env, err = loadEnvVars()
	return
}

type envVars struct {
	ACCESS_TOKEN_SECRET    string `mapstructure:"ACCESS_TOKEN_SECRET"`
	ACCESS_TOKEN_LIFETIME  string `mapstructure:"ACCESS_TOKEN_LIFETIME"`
	REFRESH_TOKEN_SECRET   string `mapstructure:"REFRESH_TOKEN_SECRET"`
	REFRESH_TOKEN_LIFETIME string `mapstructure:"REFRESH_TOKEN_LIFETIME"`
	REDIS_ADDRESS          string `mapstructure:"REDIS_ADDRESS"`
	PORT                   string `mapstructure:"PORT"`
}

func loadEnvVars() (env envVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading env file", err)
		return
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error mapping env file", err)
		return
	}

	if env.REDIS_ADDRESS == "" {
		err = errors.New("REDIS_ADDRESS is required")
		return
	}
	return
}
