package decorator

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	pizzaPrice := t.Pizza.GetPrice()
	return pizzaPrice + 7
}
