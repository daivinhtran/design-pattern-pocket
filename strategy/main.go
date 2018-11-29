package main

import (
	"design-pattern-pocket/strategy/models"
)

func main() {
	mallard := models.NewMallardDuck()
	mallard.Duck.PerformFly()
	mallard.Duck.PerformQuack()

	model := models.NewModelDuck()
	model.Duck.PerformFly()
	model.Duck.PerformQuack()
}
