package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"weather-api/internal/config"
	"weather-api/internal/service"
)

var (
	rateLimit      = 2
	rateLimitStore = make(map[string][]time.Time)
	mu             sync.Mutex
)

func isRateLimited(ip string) bool {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	times := rateLimitStore[ip]

	freshTimes := []time.Time{}
	for _, t := range times {
		if now.Sub(t) < time.Minute {
			freshTimes = append(freshTimes, t)
		}
	}

	if len(freshTimes) >= rateLimit {
		return true
	}

	freshTimes = append(freshTimes, now)
	rateLimitStore[ip] = freshTimes
	return false
}

func WeatherHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if isRateLimited(ip) {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

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
