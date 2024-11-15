package location

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/willychavez/GoWeatherLookup/app/internal/domain"
)

type LocationRepositoryImpl struct {
	client *http.Client
}

func NewLocationRepository(client *http.Client) domain.LocationRespository {
	return &LocationRepositoryImpl{client: client}
}

func (l *LocationRepositoryImpl) FindCityByZipCode(zipCode string) (string, error) {
	resp, err := l.client.Get("https://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to find city by zip code")
	}
	defer resp.Body.Close()

	Body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(Body, &result)

	if city, ok := result["localidade"].(string); ok {
		return city, nil
	}
	return "", errors.New("city not found")
}
