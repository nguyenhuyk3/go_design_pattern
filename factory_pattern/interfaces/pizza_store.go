package interfaces

type PizzaStore interface {
	CreatePizza(pizzaType string) IPizza
	OrderPizza(pizzaType string) IPizza
}

func OrderPizza(store PizzaStore, pizzaType string) IPizza {
	pizza := store.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
