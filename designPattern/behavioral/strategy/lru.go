package main

import "fmt"

type Lru struct{}

func (l *Lru) evic(c *cache) {
	fmt.Println("lru evic..")
}
