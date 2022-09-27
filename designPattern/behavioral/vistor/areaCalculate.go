package main

import "fmt"

type areaCalculate struct {
	area int
}

func NewAreaCaculate() Vistor {
	return &areaCalculate{}
}

func (a *areaCalculate) visitForCircle(c *circle) {
	fmt.Println("caculate circle area")
}

func (a *areaCalculate) visitForRectangle(r *rectangle) {
	fmt.Println("caculate rectangle area")
}

func (a *areaCalculate) visitForSquare(s *square) {
	fmt.Println("caculate  square area")
}
