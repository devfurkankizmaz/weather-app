package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	SERVER_PORT         string `mapstructure:"SERVER_PORT"`
	REDIS_ADDRESS       string `mapstructure:"REDIS_ADDRESS"`
	REDIS_PASS          string `mapstructure:"REDIS_PASS"`
	REDIS_DB            int    `mapstructure:"REDIS_DB"`
	REDIS_EXPIRY_MIN    int    `mapstructure:"REDIS_EXPIRY_MIN"`
	API_KEY             string `mapstructure:"API_KEY"`
	API_URL             string `mapstructure:"API_URL"`
	CONTEXT_TIMEOUT_SEC int    `mapstructure:"CONTEXT_TIMEOUT_SEC"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile("config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file config.yml : ", err)
		return &env, err
	}
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return &env, err
	}
	return &env, nil
}
