package models

// MallardDuck is the duck tat can fly with wings
type MallardDuck struct {
	Duck
}

// NewMallardDuck returns a new MallardDuck
func NewMallardDuck() *MallardDuck {
	mallard := &MallardDuck{}
	mallard.flyer = new(FlyWithWings) // inserts wings into MallwardDuck
	mallard.quacker = new(Quack)
	return mallard
}
