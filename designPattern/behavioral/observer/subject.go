package main

type subject interface {
	register(id int)
	deregister(id int)
	notifyAll()
}
