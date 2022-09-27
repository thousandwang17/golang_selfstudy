package main

import "fmt"

type Fifo struct{}

func (f Fifo) evic(c *cache) {
	fmt.Println("fifo evic..")
}
