# 🌦️ Clima Go API

## Description
This is a Go (Golang) API to retrieve weather data based on Azure Maps. It provides endpoints to fetch current weather information and future forecasts.

## Features

- **Current Weather Data Retrieval**: The API offers an endpoint to fetch the latest weather data for a specific location.
- **Weather Forecast**: In addition to current weather data, the API also provides weather forecasts for the upcoming days.
- **Support for Multiple Data Sources**: The API is designed to be extensible and can be easily integrated with different weather service providers.

## Endpoints

### 1. Get Current Weather Data
```
GET /api/v1//current-climate/{city}
```
Returns the current weather data for the specified latitude and longitude.

**URL Path Parameters:**
- `city`: City to be analyzed 

### 2. Get Weather Forecast
```
GET /api/v1/forecast-climate/{city}/{days}
```
Returns the weather forecast for the upcoming days for the specified latitude and longitude.

**Query Parameters:**
- `city`: City to be analyzed
- `days` (optional): Number of days for which you want the weather forecast. Default is 7 days.

## Installation and Usage

1. **Installation**:
   ```shell
   git clone github.com/your-username/weather-api](https://github.com/JoaoFel1pe/clima-go-api
   ```

2. **Execution**:
   ```shell
   go run main.go
   ```

## Configuration

Before running the API, you need to configure the API keys of the Azure Maps provider. The keys can be set in the `.env` file.

```plaintext
AZURE_API_KEY=YourAPIKey1
```

## Contributing

Contributions are welcome! If you'd like to improve this API, feel free to submit pull requests.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
