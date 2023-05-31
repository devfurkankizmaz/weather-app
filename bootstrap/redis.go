package bootstrap

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisCache(env *Env) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.CONTEXT_TIMEOUT_SEC)*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_ADDRESS,
		Password: env.REDIS_PASS,
		DB:       env.REDIS_DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err.Error())
		return client, err
	}

	return client, nil
}
