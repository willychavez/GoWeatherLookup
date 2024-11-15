package location_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/location"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/mock"
)

func TestFindCityByZipCode(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"localidade": "São Paulo"}`)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := location.NewLocationRepository(httpClient)
	city, err := repo.FindCityByZipCode("01001000")

	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", city)
}

func TestFindCityByZipCodeError(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader(`{"erro": true}`)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := location.NewLocationRepository(httpClient)
	city, err := repo.FindCityByZipCode("01001000")

	assert.Error(t, err)
	assert.Equal(t, "", city)
}

func TestFindCityByZipCodeErrorRequest(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Erro na requisição")
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := location.NewLocationRepository(httpClient)
	city, err := repo.FindCityByZipCode("01001000")

	assert.Error(t, err)
	assert.Equal(t, "", city)
}

func TestFindCityByZipCodeInvalidResponse(t *testing.T) {
	mockClient := &mock.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"localidade": "São Paulo"`)),
			}, nil
		},
	}

	httpClient := &http.Client{
		Transport: mockClient,
	}
	repo := location.NewLocationRepository(httpClient)
	city, err := repo.FindCityByZipCode("01001000")

	assert.Error(t, err)
	assert.Equal(t, "", city)
}
