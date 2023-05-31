package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/weather-app/bootstrap"
	"github.com/devfurkankizmaz/weather-app/model"
	"github.com/devfurkankizmaz/weather-app/service"
	"github.com/devfurkankizmaz/weather-app/types"
	"github.com/devfurkankizmaz/weather-app/utils"
	"github.com/labstack/echo/v4"
)

type WeatherHandler struct {
	weatherService service.WeatherService
	env            *bootstrap.Env
}

func NewWeatherHandler(weatherService service.WeatherService, env *bootstrap.Env) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
		env:            env,
	}
}

// Weather godoc
// @Summary      Show weather info
// @Description  Get weather infos by given city name in query
// @Tags         weather
// @Accept       json
// @Produce      json
// @Param        city    query     string  true  "Weather search by city name"  Format(city)
// @Success      200  {object}  types.StoreData
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router       /weather [get]
func (h *WeatherHandler) Weather(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.env.CONTEXT_TIMEOUT_SEC)*time.Second)
	defer cancel()
	var weather *model.Weather

	q := c.QueryParam("city")

	response, err := h.weatherService.GetWeatherByCity(ctx, q)

	// When the Redis key expires, It will be skip err and the data is re-stored in Redis.
	// This prevents continuous API requests with the previously stored key in Redis.
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": response})
	}

	apiKey := h.env.API_KEY
	apiAddress := h.env.API_URL

	apiUrl := fmt.Sprintf("http://%s?key=%s&q=%s", apiAddress, apiKey, q)

	body, err := utils.ApiCall(apiUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
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
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "Wrong query"})
	}
	err = h.weatherService.CreateWeather(ctx, q, store)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": store})
}
