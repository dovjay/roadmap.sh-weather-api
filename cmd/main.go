package main

import (
	"log"
	"net/http"

	"weather-api/internal/config"
	"weather-api/internal/handler"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/api/weather", handler.WeatherHandler(cfg))

	log.Println("Server is running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
