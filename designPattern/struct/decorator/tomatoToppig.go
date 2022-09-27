package main

type tomatoTopping struct{}

func AddTomatoTopping() IPizza {
	return &tomatoTopping{}
}

func (c *tomatoTopping) GetPrice(add func() int) func() int {
	return func() int {
		return 7 + add()
	}
}
