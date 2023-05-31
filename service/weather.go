package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/devfurkankizmaz/weather-app/repository"
	"github.com/devfurkankizmaz/weather-app/types"
)

type weatherService struct {
	weatherRepository repository.WeatherRepository
	contextTimeout    time.Duration
}

type WeatherService interface {
	CreateWeather(ctx context.Context, city string, weather *types.StoreData) error
	GetWeatherByCity(ctx context.Context, city string) (types.StoreData, error)
}

func NewWeatherService(weatherRepository repository.WeatherRepository, timeout time.Duration) WeatherService {
	return &weatherService{
		weatherRepository: weatherRepository,
		contextTimeout:    timeout,
	}
}

func (s *weatherService) CreateWeather(ctx context.Context, city string, weather *types.StoreData) error {
	key := fmt.Sprintf("weather:%s", city)
	w, err := json.Marshal(weather)
	if err != nil {
		return err
	}
	err = s.weatherRepository.CreateWeather(ctx, key, w)
	if err != nil {
		return err
	}
	return nil
}

func (s *weatherService) GetWeatherByCity(ctx context.Context, city string) (types.StoreData, error) {
	key := fmt.Sprintf("weather:%s", city)
	data := types.StoreData{}
	cachedWeather, err := s.weatherRepository.GetWeatherByCity(ctx, key)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(cachedWeather, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
