package main

type sheap interface {
	getType() string
	accept(Vistor)
}
