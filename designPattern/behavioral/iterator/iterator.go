package main

type iterator interface {
	HasNext() bool
	getNext() *user
}
