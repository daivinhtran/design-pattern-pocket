package models

import (
	"fmt"
	"math"
)

// StatisticsDisplay implements Observer and DisplayElement interface
type StatisticsDisplay struct {
	maxTemperature float32
	minTemperature float32
	weatherData    Subject
}

// NewStatisticsDisplay returns a new CurrentConditionDisplay object
func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	sd := &StatisticsDisplay{
		weatherData:    weatherData,
		maxTemperature: -math.MaxFloat32,
		minTemperature: math.MaxFloat32,
	}
	weatherData.RegisterObserver(sd)
	return sd
}

// Update is called by the Subject (Notifier) object when there is a change
func (sd *StatisticsDisplay) Update(temperature float32, humidity float32, pressure float32) {
	sd.maxTemperature = float32(math.Max(float64(temperature), float64(sd.maxTemperature)))
	sd.minTemperature = float32(math.Min(float64(temperature), float64(sd.minTemperature)))
	sd.Display()
}

// Display is called when there is a change being triggered from the Subject object
func (sd *StatisticsDisplay) Display() {
	fmt.Printf("Max/Min temperature = %.2f/%.2f\n", sd.maxTemperature, sd.minTemperature)
}
