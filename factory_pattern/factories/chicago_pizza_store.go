package factories

import (
	"design_pattern/factory_pattern/interfaces"
	"design_pattern/factory_pattern/products"
)

type ChicagoPizzaStore struct{}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) interfaces.IPizza {
	if pizzaType == "cheese" {
		return products.NewChicagoStyleCheesePizza()
	}
	return nil
}

func (s *ChicagoPizzaStore) OrderPizza(pizzaType string) interfaces.IPizza {
	// * OrderPizza will be called at the factory
	return interfaces.OrderPizza(s, pizzaType)
}
