package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type WeatherResponse struct {
	ResolvedAddress string `json:"resolvedAddress"`
	Timezone        string `json:"timezone"`
	Days            []Day  `json:"days"`
}

type Day struct {
	Datetime    string  `json:"datetime"`
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
	Conditions  string  `json:"conditions"`
}

var rdb *redis.Client

func SetRedisClient(client *redis.Client) {
	rdb = client
}

func GetWeather(apiKey, location string) (WeatherResponse, error) {
	var result WeatherResponse
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", location, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return result, fmt.Errorf("failed to fetch weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

func GetWeatherWithCache(apiKey, location string) (WeatherResponse, error) {
	var cached WeatherResponse
	ctx := context.Background()

	val, err := rdb.Get(ctx, location).Result()
	if err == nil {
		err := json.Unmarshal([]byte(val), &cached)
		if err == nil {
			return cached, nil
		}
	}

	fresh, err := GetWeather(apiKey, location)
	if err != nil {
		return fresh, err
	}

	jsonData, err := json.Marshal(fresh)
	if err == nil {
		rdb.Set(ctx, location, jsonData, 10*time.Minute)
	}

	return fresh, nil
}
