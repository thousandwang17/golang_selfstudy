package main

import "fmt"

type red struct{}

func NewRed() icolor {
	return &red{}
}

func (r *red) printColor() {
	fmt.Println("i'm red")
}
