package models

// Duck class respresents all ducks in the universe
type Duck struct {
	flyer   Flyer
	quacker Quacker
}

// PerformFly ...
func (d *Duck) PerformFly() {
	d.flyer.fly()
}

// PerformQuack ...
func (d *Duck) PerformQuack() {
	d.quacker.quack()
}
