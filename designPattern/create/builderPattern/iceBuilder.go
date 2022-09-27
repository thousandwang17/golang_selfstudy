package main

type iceBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIcebuilder() *iceBuilder {
	return &iceBuilder{}
}

func (i *iceBuilder) setDoorType() {
	i.doorType = "water"
}

func (i *iceBuilder) setWindowType() {
	i.windowType = "water"
}

func (i *iceBuilder) setnumFloor() {
	i.floor = 1
}

func (i *iceBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
