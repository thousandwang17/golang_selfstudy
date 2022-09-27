package main

type tver interface {
	on() error
	off() error
}
