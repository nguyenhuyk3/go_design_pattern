package ducks

import (
	"design_pattern/stratergy_pattern/behaviors"
	"fmt"
)

type Duck struct {
	FlyBehavior   behaviors.IFlyBehavior
	QuackBehavior behaviors.IQuackBehavior
}

func (d *Duck) SetFlyBehavior(fb behaviors.IFlyBehavior) {
	d.FlyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb behaviors.IQuackBehavior) {
	d.QuackBehavior = qb
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.QuackBehavior.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("Tôi đang bơi!!")
}
