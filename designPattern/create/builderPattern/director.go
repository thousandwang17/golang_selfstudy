package main

type director struct {
	builder ibuilder
}

func newDirector(b ibuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b ibuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setnumFloor()

	return d.builder.getHouse()
}
