package implementations

import "fmt"

type ConcreteObserver struct {
	Name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s nhận được thông báo: %s\n", o.Name, message)
}
