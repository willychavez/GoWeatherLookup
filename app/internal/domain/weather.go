package domain

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type LocationRespository interface {
	FindCityByZipCode(zipCode string) (string, error)
}

type WeatherRepository interface {
	GetCurrentTemperature(city string) (float64, error)
}
