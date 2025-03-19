package ducks

import (
	"design_pattern/stratergy_pattern/implementations/fly"
	"design_pattern/stratergy_pattern/implementations/quack"
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
