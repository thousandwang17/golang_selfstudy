package main

type ibuilder interface {
	setWindowType()
	setnumFloor()
	setDoorType()
	getHouse() house
}

func getBuilder(b string) ibuilder {
	switch b {
	case "iceHouse":
		return &iceBuilder{}
	case "normalHouse":
		return &normalBuilder{}
	}
	return nil
}
