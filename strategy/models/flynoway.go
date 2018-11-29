package models

import "fmt"

// FlyNoWay represents the implementation of flying for all ducks that can't fly
type FlyNoWay struct{}

func (fnw *FlyNoWay) fly() {
	fmt.Println("I can't fly")
}
