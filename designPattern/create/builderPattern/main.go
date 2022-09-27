package main

import "fmt"

func main() {
	// ice house
	builder := getBuilder("iceHouse")
	director := newDirector(builder)
	house := director.buildHouse()

	fmt.Printf(" doorType : %v \n", house.doorType)
	fmt.Printf(" windowType : %v \n", house.windowType)
	fmt.Printf(" floor : %v \n", house.floor)

	// normalHouse
	builder = getBuilder("normalHouse")
	director.setBuilder(builder)
	house = director.buildHouse()

	fmt.Printf(" doorType : %v \n", house.doorType)
	fmt.Printf(" windowType : %v \n", house.windowType)
	fmt.Printf(" floor : %v \n", house.floor)
}
