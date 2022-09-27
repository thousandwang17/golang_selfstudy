package main

type Vistor interface {
	visitForCircle(*circle)
	visitForSquare(*square)
	visitForRectangle(*rectangle)
}
