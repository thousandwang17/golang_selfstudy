package main

type veggieMania struct{}

func AddVeggieMania() IPizza {
	return &veggieMania{}
}

func (c *veggieMania) GetPrice(add func() int) func() int {
	return func() int {
		return 10 + add()
	}
}
