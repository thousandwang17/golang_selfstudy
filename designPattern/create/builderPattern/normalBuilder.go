package main

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden door"
}

func (b *normalBuilder) setnumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
