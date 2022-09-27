package main

type patient struct {
	name string
}

func NewPatient(n string) *patient {
	return &patient{name: n}
}
