package main

type TerroristDress struct {
	color string
}

func NewTerroristDress() IDress {
	return &TerroristDress{
		color: "red",
	}
}

func (p *TerroristDress) getColor() string {
	return p.color
}
