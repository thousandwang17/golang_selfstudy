package main

type circle struct {
	radius int
}

func NewCircle(r int) sheap {
	return &circle{
		radius: r,
	}
}

func (c *circle) getType() string {
	return "circle"
}

func (c *circle) accept(v Vistor) {
	v.visitForCircle(c)
}
