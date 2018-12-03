package models

// Observer is an interface for observers
type Observer interface {
	Update(float32, float32, float32)
}
