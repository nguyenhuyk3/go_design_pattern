package observerpattern

import (
	"design_pattern/observer_pattern/implementations"
	"fmt"
)

func Run() {
	subj := &implementations.ConcreteSubject{}

	obs1 := &implementations.ConcreteObserver{Name: "Observer 1"}
	obs2 := &implementations.ConcreteObserver{Name: "Observer 2"}

	subj.Register(obs1)
	subj.Register(obs2)

	fmt.Println("Thông báo lần 1:")
	subj.Notify("Hello Observers!")

	subj.Deregister(obs1)

	fmt.Println("Thông báo lần 2:")
	subj.Notify("Hello again!")
}
