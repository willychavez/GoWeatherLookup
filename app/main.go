package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/api"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/location"
	"github.com/willychavez/GoWeatherLookup/app/internal/inferfaces/repository/weather"
	"github.com/willychavez/GoWeatherLookup/app/internal/infrastructure/httpclient"
	"github.com/willychavez/GoWeatherLookup/app/internal/usecase"
)

func main() {
	httpClient := httpclient.NewHttpClient()

	locationRepo := location.NewLocationRepository(httpClient)
	weatherRepo := weather.NewWeatherRepository(httpClient)
	weatherUseCase := usecase.WeatherUseCase{
		LocationRepo: locationRepo,
		WeatherRepo:  weatherRepo,
	}

	weatherHandler := api.WeatherHandler{
		UseCase: &weatherUseCase,
	}

	r := mux.NewRouter()
	r.HandleFunc("/weather", weatherHandler.GetWeather).Methods("GET")
	r.HandleFunc("/healthz", api.HealthHandler).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
