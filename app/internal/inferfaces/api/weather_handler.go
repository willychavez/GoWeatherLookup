package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/willychavez/GoWeatherLookup/app/internal/usecase"
)

type WeatherHandler struct {
	UseCase *usecase.WeatherUseCase
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipCode := r.URL.Query().Get("zipcode")
	log.Println("Get weather by zipcode: ", zipCode)
	if len(zipCode) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.UseCase.GetWeatherByZipCode(zipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
