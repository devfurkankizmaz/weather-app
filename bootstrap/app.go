package bootstrap

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Env    *Env
	Client *redis.Client
	Logger *logrus.Logger
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	app.Env = env

	client, err := NewRedisCache(app.Env)
	if err != nil {
		log.Fatal(err)
	}
	app.Client = client

	logger := InitLogger()
	app.Logger = logger
	return *app
}
