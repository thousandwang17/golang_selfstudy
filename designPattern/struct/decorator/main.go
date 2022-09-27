package main

import "fmt"

func main() {
	pizza := NewPizza()

	pizza.AddTopping(AddChessTopping())
	pizza.AddTopping(AddTomatoTopping())
	pizza.AddTopping(AddVeggieMania())

	fmt.Printf("price %d ", pizza.GetPrice())
}
