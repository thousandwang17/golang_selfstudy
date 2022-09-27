package main

type rectangle struct {
	l, b int
}

func NewRectangle(l, b int) sheap {
	return &rectangle{
		l: l,
		b: b,
	}
}

func (r *rectangle) getType() string {
	return "rectangle"
}

func (r *rectangle) accept(v Vistor) {
	v.visitForRectangle(r)
}
