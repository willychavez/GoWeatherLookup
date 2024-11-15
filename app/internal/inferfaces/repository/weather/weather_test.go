package weather_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/mock"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/weather"
)

func TestGetCurrentTemperature(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"current": {"temp_c": 25.0}}`)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := weather.NewWeatherRepository(httpClient)
	temp, err := repo.GetCurrentTemperature("TestCity")

	assert.NoError(t, err)
	assert.Equal(t, 25.0, temp)
}

func TestGetCurrentTemperature_Failure(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("failed to fetch weather data")
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := weather.NewWeatherRepository(httpClient)
	temp, err := repo.GetCurrentTemperature("TestCity")

	assert.Error(t, err)
	assert.Equal(t, 0.0, temp)
}

func TestGetCurrentTemperature_InvalidStatusCode(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader(``)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := weather.NewWeatherRepository(httpClient)
	temp, err := repo.GetCurrentTemperature("TestCity")

	assert.Error(t, err)
	assert.Equal(t, 0.0, temp)
}

func TestGetCurrentTemperature_TemperatureNotFound(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"current": {}}`)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := weather.NewWeatherRepository(httpClient)
	temp, err := repo.GetCurrentTemperature("TestCity")

	assert.Error(t, err)
	assert.Equal(t, 0.0, temp)
}
