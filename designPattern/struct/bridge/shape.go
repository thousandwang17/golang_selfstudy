package main

type iShape interface {
	GetArea() float32
	GetPerimeter() float32
	GetColor()
	SetColor(icolor)
}
