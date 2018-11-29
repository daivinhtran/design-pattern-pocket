package models

import "fmt"

// FlyWithWings represents the implementation of flying for all ducks that have wings
type FlyWithWings struct{}

// implements Flyer interface
func (fww FlyWithWings) fly() {
	fmt.Println("Flying with Wings")
}
