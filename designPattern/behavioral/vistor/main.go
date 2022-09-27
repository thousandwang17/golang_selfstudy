package main

func main() {
	square := NewSquare(5)
	circle := NewCircle(3)
	rectangle := NewRectangle(1, 2)

	areaCalculate := NewAreaCaculate()

	square.accept(areaCalculate)
	circle.accept(areaCalculate)
	rectangle.accept(areaCalculate)

	// fmt.Println()
	middleCoordinate := NewMiddleCoordinate()

	square.accept(middleCoordinate)
	circle.accept(middleCoordinate)
	rectangle.accept(middleCoordinate)
}
