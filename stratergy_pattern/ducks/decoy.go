package ducks

import (
	"design_pattern/stratergy_pattern/implementations/fly"
	"design_pattern/stratergy_pattern/implementations/quack"
)

type DecoyDuck struct {
	Duck
}

func NewDecoyDuck() *DecoyDuck {
	return &DecoyDuck{
		Duck{
			FlyBehavior:   &fly.FlyNoWay{},
			QuackBehavior: &quack.MuteQuack{},
		},
	}
}
