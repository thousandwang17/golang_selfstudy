package main

import "fmt"

type middleCoordinate struct {
	x, y int
}

func NewMiddleCoordinate() Vistor {
	return &middleCoordinate{}
}

func (a *middleCoordinate) visitForCircle(c *circle) {
	fmt.Println("finding circle middleCoordinate")
}

func (a *middleCoordinate) visitForRectangle(r *rectangle) {
	fmt.Println("finding rectangle middleCoordinate")
}

func (a *middleCoordinate) visitForSquare(s *square) {
	fmt.Println("finding  square middleCoordinate")
}
