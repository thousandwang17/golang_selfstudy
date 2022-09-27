package main

import "fmt"

func main() {
	red := NewRed()
	blue := NewBlue()

	circle := NewCircle(3.3, red)
	square := NewSquare(4, blue)

	// circle
	fmt.Printf("circle Area %v, permiter %v \n", circle.GetArea(), circle.GetPerimeter())
	circle.GetColor()
	circle.SetColor(blue)
	circle.GetColor()

	// square
	fmt.Printf("circle Area %v, permiter %v \n", square.GetArea(), square.GetPerimeter())
	square.GetColor()
	square.SetColor(red)
	square.GetColor()
}
