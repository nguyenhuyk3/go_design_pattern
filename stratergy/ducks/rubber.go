package ducks

import (
	"design_pattern/stratergy/implementations/fly"
	"design_pattern/stratergy/implementations/quack"
)

type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{
		Duck{
			FlyBehavior:   &fly.FlyNoWay{},
			QuackBehavior: &quack.Squeak{},
		},
	}
}
