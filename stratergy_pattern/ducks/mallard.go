package ducks

import (
	"design_pattern/stratergy_pattern/implementations/fly"
	"design_pattern/stratergy_pattern/implementations/quack"
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
