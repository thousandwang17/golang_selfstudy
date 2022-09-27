package main

type square struct {
	side int
}

func NewSquare(side int) sheap {
	return &square{
		side: side,
	}
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(v Vistor) {
	v.visitForSquare(s)
}
