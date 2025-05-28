# Weather API (Go)

This project provides a simple Go-based weather API that fetches weather data using the [Visual Crossing Weather API](https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/).

Part of Roadmap.sh project [Weather API](https://roadmap.sh/projects/weather-api-wrapper-service)

## Features

- 🌦️ Get current weather for a given location
- 🚫 Rate limiting (10 requests/minute per IP)
- ⚡ Redis-based caching (10-minute expiry)
- 🔐 Configurable via `.env`

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/dovjay/roadmap.sh-weather-api
cd weather-api
```

### 2. Create `.env` file

```env
VISUAL_CROSSING_API_KEY=your_api_key_here
PORT=8080
```

### 3. Install dependencies

```bash
go mod tidy
```

### 5. Start the server

```bash
go run cmd/main.go
```

## API Usage

### Endpoint

```
GET /api/weather?location=Jakarta
```

### Query Parameters

- `location` (required): Name of the city or location or long/lat.

### Example Response

```json
{
  "resolvedAddress": "Jakarta, Indonesia",
  "timezone": "Asia/Jakarta",
  "days": [
    {
      "datetime": "2024-05-28",
      "temp": 31.5,
      "description": "Hot and humid",
      "conditions": "Clear"
    }
  ]
}
```
