package quack

import "fmt"

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Chít chít!!")
}
