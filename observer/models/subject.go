package models

// Subject is interface for notifier
type Subject interface {
	RegisterObserver(Observer)
	// RemoveObserver(Observer)
	NotifyObservers()
}
