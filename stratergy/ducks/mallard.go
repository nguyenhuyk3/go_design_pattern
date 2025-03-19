package ducks

import (
	"design_pattern/stratergy/implementations/fly"
	"design_pattern/stratergy/implementations/quack"
)

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Duck{
			FlyBehavior:   &fly.FlyWithWings{},
			QuackBehavior: &quack.Squeak{},
		},
	}
}
