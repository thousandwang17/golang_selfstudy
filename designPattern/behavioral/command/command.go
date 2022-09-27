package main

type commander interface {
	execute() error
}
