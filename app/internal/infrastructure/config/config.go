package config

import "os"

func GetWeatherApiKey() string {
	return os.Getenv("WEATHER_API_KEY")
}
