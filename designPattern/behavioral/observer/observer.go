package main

type observer interface {
	update(string)
	getId() int
	setId(int)
}
