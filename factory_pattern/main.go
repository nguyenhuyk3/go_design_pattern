package factorypattern

import (
	"design_pattern/factory_pattern/factories"
	"fmt"
)

func Run() {
	nyStore := &factories.NYPizzaStore{}
	nyPizza := nyStore.OrderPizza("cheese")
	fmt.Println("Đã đặt:", nyPizza.GetName())

	chicagoStore := &factories.ChicagoPizzaStore{}
	chicagoPizza := chicagoStore.OrderPizza("cheese")
	fmt.Println("Đã đặt:", chicagoPizza.GetName())
}
