package main

import (
	"log"
	"net/http"

	"weather-api/internal/config"
	"weather-api/internal/handler"
	"weather-api/internal/service"

	"github.com/redis/go-redis/v9"
)

func main() {
	cfg := config.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	service.SetRedisClient(rdb)

	http.HandleFunc("/api/weather", handler.WeatherHandler(cfg))

	log.Println("Server is running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
