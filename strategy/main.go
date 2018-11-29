package main

import (
	"design-pattern-pocket/strategy/models"
)

func main() {
	mallard := models.NewMallardDuck()
	mallard.PerformFly()
	mallard.PerformQuack()

	model := models.NewModelDuck()
	model.PerformFly()
	model.PerformQuack()
}
