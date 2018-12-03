package main

import (
	"design-pattern-pocket/observer/models"
)

func main() {
	weatherData := models.NewWeatherData()
	models.NewCurrentConditionDisplay(weatherData)
	models.NewStatisticsDisplay(weatherData)
	weatherData.SetMeasurements(float32(32), float32(32), float32(32))
	weatherData.SetMeasurements(float32(66), float32(32), float32(32))
	weatherData.SetMeasurements(float32(10), float32(32), float32(32))
}
