package main

import "math"

type circle struct {
	radius float32
	color  icolor
}

func NewCircle(radius float32, c icolor) iShape {
	return &circle{
		radius: radius,
		color:  c,
	}
}

func (c *circle) GetArea() float32 {
	return c.radius * c.radius * math.Pi
}

func (c *circle) GetPerimeter() float32 {
	return c.radius * 2 * math.Pi
}

func (c *circle) GetColor() {
	if c.color != nil {
		c.color.printColor()
	}
}

func (c *circle) SetColor(n icolor) {
	c.color = n
}
