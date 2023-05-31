package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/devfurkankizmaz/weather-app/model"
	"github.com/devfurkankizmaz/weather-app/repository"
	"github.com/devfurkankizmaz/weather-app/types"
	"github.com/devfurkankizmaz/weather-app/utils"
)

type weatherService struct {
	weatherRepository repository.WeatherRepository
	contextTimeout    time.Duration
}

type WeatherService interface {
	CreateWeather(ctx context.Context, url *types.Api) (*types.StoreData, error)
	GetWeatherByCity(ctx context.Context, city string) (types.StoreData, error)
}

func NewWeatherService(weatherRepository repository.WeatherRepository, timeout time.Duration) WeatherService {
	return &weatherService{
		weatherRepository: weatherRepository,
		contextTimeout:    timeout,
	}
}

func (s *weatherService) CreateWeather(ctx context.Context, url *types.Api) (*types.StoreData, error) {
	var weather *model.Weather

	apiUrl := fmt.Sprintf("http://%s?key=%s&q=%s", url.Url, url.ApiKey, url.City)

	body, err := utils.ApiCall(apiUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	store := &types.StoreData{
		Name:        weather.Location.Name,
		Region:      weather.Location.Region,
		Country:     weather.Location.Country,
		Latitude:    weather.Location.Lat,
		Longitude:   weather.Location.Lon,
		LocalTime:   weather.Location.LocalTime,
		TempC:       weather.Current.TempC,
		TempF:       weather.Current.TempF,
		LastUpdated: weather.Current.LastUpdated,
	}
	if store.Name == "" {
		return nil, fmt.Errorf("wrong query city name")
	}
	lowerName := strings.ToLower(store.Name)
	lowerCity := strings.ToLower(url.City)

	newCity := strings.Replace(lowerCity, "%", " ", -1)

	if lowerName != newCity {
		return nil, fmt.Errorf("wrong query city name")
	}

	key := fmt.Sprintf("weather:%s", strings.ToLower(url.City))
	w, err := json.Marshal(store)
	if err != nil {
		return nil, err
	}
	err = s.weatherRepository.CreateWeather(ctx, key, w)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *weatherService) GetWeatherByCity(ctx context.Context, city string) (types.StoreData, error) {
	key := fmt.Sprintf("weather:%s", strings.ToLower(city))
	var data types.StoreData
	cachedWeather, err := s.weatherRepository.GetWeatherByCity(ctx, key)
	if err != nil {
		return types.StoreData{}, err
	}
	err = json.Unmarshal(cachedWeather, &data)
	if err != nil {
		return types.StoreData{}, err
	}
	return data, nil
}
