package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type weatherRepository struct {
	client *redis.Client
	expiry time.Duration
}

type WeatherRepository interface {
	CreateWeather(ctx context.Context, key string, weather []byte) error
	GetWeatherByCity(ctx context.Context, key string) ([]byte, error)
}

func NewWeatherRepository(client *redis.Client, expiry time.Duration) WeatherRepository {
	return &weatherRepository{
		client: client,
		expiry: expiry,
	}
}

func (r *weatherRepository) CreateWeather(ctx context.Context, key string, weather []byte) error {
	err := r.client.Set(ctx, key, weather, r.expiry).Err()
	if err != nil {
		return fmt.Errorf("failed to set weather in cache: %w", err)
	}

	return nil
}

func (r *weatherRepository) GetWeatherByCity(ctx context.Context, key string) ([]byte, error) {
	cachedWeather, err := r.client.Get(ctx, key).Bytes()

	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get weather from cache: %w", err)
	}
	return cachedWeather, nil
}
