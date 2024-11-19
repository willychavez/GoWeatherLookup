package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/willychavez/GoWeatherLookup/app/internal/domain"
)

type WeatherUseCase struct {
	LocationRepo domain.LocationRespository
	WeatherRepo  domain.WeatherRepository
}

func (w *WeatherUseCase) GetWeatherByZipCode(zipCode string) (*domain.WeatherResponse, error) {
	city, err := w.LocationRepo.FindCityByZipCode(zipCode)
	if err != nil {
		return nil, errors.New("can not find zipcode")
	}

	log.Println("City found: ", city)
	tempC, err := w.WeatherRepo.GetCurrentTemperature(city)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch the current temperature for city: %s", city)
	}

	return &domain.WeatherResponse{
		TempC: tempC,
		TempF: (tempC * 1.8) + 32,
		TempK: tempC + 273,
	}, nil
}
