package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestGetWeaterByZyCode_Success(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	weatherUseCase := &WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	mockLocationRepo.On("FindCityByZipCode", "12345").Return("TestCity", nil)
	mockWeatherRepo.On("GetCurrentTemperature", "TestCity").Return(25.0, nil)

	weatherResponse, err := weatherUseCase.GetWeatherByZipCode("12345")

	assert.NoError(t, err)
	assert.NotNil(t, weatherResponse)
	assert.Equal(t, 25.0, weatherResponse.TempC)
	assert.Equal(t, 77.0, weatherResponse.TempF)
	assert.Equal(t, 298.0, weatherResponse.TempK)
}

func TestGetWeaterByZyCode_LocationRepoError(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	weatherUseCase := &WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	mockLocationRepo.On("FindCityByZipCode", "12345").Return("", errors.New("can not find zipcode"))

	weatherResponse, err := weatherUseCase.GetWeatherByZipCode("12345")

	assert.Error(t, err)
	assert.Nil(t, weatherResponse)
	assert.Equal(t, "can not find zipcode", err.Error())
}

func TestGetWeaterByZyCode_WeatherAPIError(t *testing.T) {
	mockLocationRepo := new(MockLocationRepo)
	mockWeatherRepo := new(MockWeatherRepo)
	weatherUseCase := &WeatherUseCase{
		LocationRepo: mockLocationRepo,
		WeatherRepo:  mockWeatherRepo,
	}

	mockLocationRepo.On("FindCityByZipCode", "12345").Return("TestCity", nil)
	mockWeatherRepo.On("GetCurrentTemperature", "TestCity").Return(0.0, errors.New("failed to get current temperature"))

	weatherResponse, err := weatherUseCase.GetWeatherByZipCode("12345")

	assert.Error(t, err)
	assert.Nil(t, weatherResponse)
	assert.Equal(t, "unable to fetch the current temperature for city: TestCity", err.Error())
}
