package api_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/willychavez/GoWeatherLookup/app/internal/domain"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/api"
	"github.com/willychavez/GoWeatherLookup/app/internal/usecase"
)

// Mock for LocationRepository
type MockLocationRepo struct {
	mock.Mock
}

func (m *MockLocationRepo) FindCityByZipCode(zipCode string) (string, error) {
	args := m.Called(zipCode)
	return args.String(0), args.Error(1)
}

// Mock for WeatherRepository
type MockWeatherRepo struct {
	mock.Mock
}

func (m *MockWeatherRepo) GetCurrentTemperature(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

func TestGetWeather_Success(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	useCase := &usecase.WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	mockLocationRepo.On("FindCityByZipCode", "01001000").Return("TestCity", nil)
	mockWeatherRepo.On("GetCurrentTemperature", "TestCity").Return(25.0, nil)

	expectedWeather := domain.WeatherResponse{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.0,
	}

	handler := &api.WeatherHandler{
		UseCase: useCase,
	}

	req := httptest.NewRequest("GET", "/weather?zipcode=01001000", nil)
	rec := httptest.NewRecorder()

	handler.GetWeather(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var actualWeather domain.WeatherResponse
	json.NewDecoder(res.Body).Decode(&actualWeather)
	assert.Equal(t, expectedWeather, actualWeather)

}

func TestGetWeather_InvalidZipCode(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	useCase := &usecase.WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	handler := &api.WeatherHandler{
		UseCase: useCase,
	}

	req := httptest.NewRequest("GET", "/weather?zipcode=12345", nil)
	rec := httptest.NewRecorder()

	handler.GetWeather(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestGetWeather_LocationRepoError(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	useCase := &usecase.WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	mockLocationRepo.On("FindCityByZipCode", "01001000").Return("", errors.New("can not find zipcode"))

	handler := &api.WeatherHandler{
		UseCase: useCase,
	}

	req := httptest.NewRequest("GET", "/weather?zipcode=01001000", nil)
	rec := httptest.NewRecorder()

	handler.GetWeather(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
