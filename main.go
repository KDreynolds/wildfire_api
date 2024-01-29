package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WeatherData models the weather data we want to return
type WeatherData struct {
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	WindSpeed         float64 `json:"wind_speed"`
	WindDirection     float64 `json:"wind_direction"`
	Precipitation     float64 `json:"precipitation"`
	DroughtConditions string  `json:"drought_conditions"`
	FireRisk          string  `json:"fire_risk"`
}

func main() {
	router := gin.Default()

	// Define a route to handle WeatherData requests
	router.GET("/weather", func(c *gin.Context) {
		// Placeholder for fetching and processing data
		data := WeatherData{
			// Populate with dummy data for now
			Latitude:          34.0522,
			Longitude:         -118.2437,
			Temperature:       25.0,
			Humidity:          30.0,
			WindSpeed:         5.0,
			WindDirection:     270,
			Precipitation:     0,
			DroughtConditions: "Moderate",
			FireRisk:          "High",
		}

		c.JSON(http.StatusOK, data)
	})

	// Run the server
	router.Run(":8080")
}
