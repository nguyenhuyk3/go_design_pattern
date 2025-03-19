package stratergypattern

import (
	"design_pattern/stratergy_pattern/ducks"
	"design_pattern/stratergy_pattern/implementations/fly"
	"design_pattern/stratergy_pattern/implementations/quack"
	"fmt"
)

func Run() {
	mallard := ducks.NewMallardDuck()
	fmt.Println("Mallard Duck:")
	mallard.PerformFly()
	mallard.PerformQuack()
	mallard.Swim()

	fmt.Println("\nRubber Duck:")
	rubber := ducks.NewRubberDuck()
	rubber.PerformFly()
	rubber.PerformQuack()
	rubber.Swim()

	fmt.Println("\nDecoy Duck:")
	decoy := ducks.NewDecoyDuck()
	decoy.PerformFly()
	decoy.PerformQuack()
	decoy.Swim()

	fmt.Println("\nThay đổi hành vi của Rubber Duck:")
	rubber.SetFlyBehavior(&fly.FlyWithWings{})
	rubber.SetQuackBehavior(&quack.Squeak{})
	rubber.PerformFly()
	rubber.PerformQuack()

}
