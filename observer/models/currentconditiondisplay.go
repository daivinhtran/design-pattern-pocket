package models

import "fmt"

// CurrentConditionDisplay implements Observer and DisplayElement interface
type CurrentConditionDisplay struct {
	temperature float32
	humidity    float32
	weatherData Subject
}

// NewCurrentConditionDisplay returns a new CurrentConditionDisplay object
func NewCurrentConditionDisplay(weatherData Subject) *CurrentConditionDisplay {
	ccd := &CurrentConditionDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(ccd)
	return ccd
}

// Update is called by the Subject (Notifier) object when there is a change
func (ccd *CurrentConditionDisplay) Update(temperature float32, humidity float32, pressure float32) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.Display()
}

// Display is called when there is a change being triggered from the Subject object
func (ccd *CurrentConditionDisplay) Display() {
	fmt.Printf("Current condiditions %.2fF degress and %.2f humidity\n", ccd.temperature, ccd.humidity)
}
