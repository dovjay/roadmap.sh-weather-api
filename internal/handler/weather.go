package handler

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/config"
	"weather-api/internal/service"
)

func WeatherHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		if location == "" {
			http.Error(w, "location is required", http.StatusBadRequest)
			return
		}

		data, err := service.GetWeather(cfg.APIKey, location)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
