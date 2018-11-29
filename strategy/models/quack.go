package models

import "fmt"

// Quack represents the implementation of flying for all ducks that can quack
type Quack struct{}

// implements Flyer interface
func (q Quack) quack() {
	fmt.Println("Quack quack!")
}
