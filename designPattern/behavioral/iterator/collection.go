package main

type collection interface {
	CreateIterator(t int) iterator // type
	Insert(u ...*user)
}
