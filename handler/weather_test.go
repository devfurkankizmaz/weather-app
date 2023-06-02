package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devfurkankizmaz/weather-app/bootstrap"
	"github.com/devfurkankizmaz/weather-app/handler"
	"github.com/devfurkankizmaz/weather-app/service"
	"github.com/devfurkankizmaz/weather-app/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWeatherHandler(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWeatherResp := types.StoreData{}

		mockWeatherService := &service.MockWeatherService{}

		arg1 := mock.Anything
		arg2 := mock.Anything
		mockWeatherService.On("GetWeatherByCity", arg1, arg2).Return(mockWeatherResp, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/weather?city=Istanbul", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewWeatherHandler(
			mockWeatherService,
			&bootstrap.Env{},
		)
		weatherJSON, err := json.Marshal(echo.Map{
			"data":   mockWeatherResp,
			"status": "success",
		})
		body := string(weatherJSON) + "\n"

		assert.NoError(t, err)

		if assert.NoError(t, h.Weather(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, body, rec.Body.String())
		}

	})
	t.Run("NotFound", func(t *testing.T) {
		mockWeatherService := &service.MockWeatherService{}

		arg1 := mock.Anything
		arg2 := mock.Anything
		mockWeatherService.On("GetWeatherByCity", arg1, arg2).Return(nil, fmt.Errorf(""))
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/weather", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewWeatherHandler(
			mockWeatherService,
			&bootstrap.Env{},
		)

		respErr, err := json.Marshal(echo.Map{
			"message": "city param can not empty",
			"status":  "error",
		})
		body := string(respErr) + "\n"

		assert.NoError(t, err)
		if assert.NoError(t, h.Weather(c)) {
			assert.Equal(t, body, rec.Body.String())
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}

	})
}
