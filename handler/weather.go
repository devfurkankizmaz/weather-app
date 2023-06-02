package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/weather-app/bootstrap"
	"github.com/devfurkankizmaz/weather-app/service"
	"github.com/devfurkankizmaz/weather-app/types"
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
	q := c.QueryParam("city")
	param := c.QueryParams().Get("city")
	if param == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "error", "message": "city param can not empty"})
	}

	response, err := h.weatherService.GetWeatherByCity(ctx, q)
	// When the Redis key expires, It will be skip err and the data is re-stored in Redis.
	// This prevents continuous API requests with the previously stored key in Redis.
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": response})
	}

	apiUrl := &types.Api{
		Url:    h.env.API_URL,
		City:   q,
		ApiKey: h.env.API_KEY,
	}

	store, err := h.weatherService.CreateWeather(ctx, apiUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": store})
}
