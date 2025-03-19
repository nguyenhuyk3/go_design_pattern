package factories

import (
	"design_pattern/factory_pattern/interfaces"
	"design_pattern/factory_pattern/products"
)

type NYPizzaStore struct{}

func (s *NYPizzaStore) CreatePizza(pizzaType string) interfaces.IPizza {
	if pizzaType == "cheese" {
		return products.NewNYStyleCheesePizza()
	}

	return nil
}

func (s *NYPizzaStore) OrderPizza(pizzaType string) interfaces.IPizza {
	return interfaces.OrderPizza(s, pizzaType)
}
