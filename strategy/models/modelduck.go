package models

// ModelDuck is the duck tat can fly with wings
type ModelDuck struct {
	Duck
}

// NewModelDuck returns a new MallardDuck
func NewModelDuck() *ModelDuck {
	model := &ModelDuck{}
	model.flyer = new(FlyNoWay) // inserts wings into MallwardDuck
	model.quacker = new(Quack)
	return model
}
