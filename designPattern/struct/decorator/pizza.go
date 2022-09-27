package main

type IPizza interface {
	GetPrice(func() int) func() int
}

type Pizza interface {
	AddTopping(t IPizza)
	GetPrice() int
}

type pizza struct {
	price   int
	Topping func() int
}

func NewPizza() Pizza {
	return &pizza{
		price:   100,
		Topping: func() int { return 0 },
	}
}

func (p *pizza) AddTopping(t IPizza) {
	p.Topping = t.GetPrice(p.Topping)
}

func (p *pizza) GetPrice() int {
	return p.Topping() + p.price
}
