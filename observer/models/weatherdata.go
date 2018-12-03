package models

// WeatherData implements Subject interface and holds the states of the app
type WeatherData struct {
	temperature float32
	humidity    float32
	pressure    float32
	observers   []Observer
}

// NewWeatherData returns a new WeatherData object
func NewWeatherData() *WeatherData {
	wd := &WeatherData{
		observers: []Observer{},
	}

	return wd
}

// RegisterObserver subscribes an observer
func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

// NotifyObservers update the states of all obeservers
func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

// SetMeasurements is called when there is a changed on the current condition
func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float32) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}

func (wd *WeatherData) measurementsChanged() {
	wd.NotifyObservers()
}
