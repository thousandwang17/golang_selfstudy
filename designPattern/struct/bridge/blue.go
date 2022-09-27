package main

import "fmt"

type blue struct{}

func NewBlue() icolor {
	return &blue{}
}

func (r *blue) printColor() {
	fmt.Println("i'm blue")
}
