package main

import "fmt"

type Lfu struct{}

func (l *Lfu) evic(c *cache) {
	fmt.Println("lfu evic..")
}
