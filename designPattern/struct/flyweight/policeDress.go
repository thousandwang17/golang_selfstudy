package main

type policeDress struct {
	color string
}

func NewPoliceDress() IDress {
	return &policeDress{
		color: "green",
	}
}

func (p *policeDress) getColor() string {
	return p.color
}
