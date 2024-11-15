package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/willychavez/GoWeatherLookup/app/internal/domain"
	"github.com/willychavez/GoWeatherLookup/app/internal/infrastructure/config"
)

type weatherRepositoryImpl struct {
	client *http.Client
}

func NewWeatherRepository(client *http.Client) domain.WeatherRepository {
	return &weatherRepositoryImpl{client: client}
}

func (repo *weatherRepositoryImpl) GetCurrentTemperature(city string) (float64, error) {
	apiKey := config.GetWeatherApiKey()
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := repo.client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0, errors.New("failed to fetch weather data")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if current, ok := result["current"].(map[string]interface{}); ok {
		if tempC, ok := current["temp_c"].(float64); ok {
			return tempC, nil
		}
	}
	return 0, errors.New("temperature not found")
}
