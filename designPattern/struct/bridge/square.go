package main

type square struct {
	size  float32
	color icolor
}

func NewSquare(size float32, c icolor) iShape {
	return &square{
		size:  size,
		color: c,
	}
}

func (c *square) GetArea() float32 {
	return c.size * c.size
}

func (c *square) GetPerimeter() float32 {
	return c.size * 4
}

func (c *square) GetColor() {
	if c.color != nil {
		c.color.printColor()
	}
}

func (c *square) SetColor(n icolor) {
	c.color = n
}
