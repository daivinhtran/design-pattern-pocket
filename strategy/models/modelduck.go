package models

// ModelDuck is the duck tat can fly can't fly but quack
type ModelDuck struct {
	Duck
}

// NewModelDuck returns a new MallardDuck
func NewModelDuck() *ModelDuck {
	model := &ModelDuck{}
	model.flyer = new(FlyNoWay) // make sure ModelDuck has no wings
	model.quacker = new(Quack)  // make sure it can quack
	return model
}
