package main

import "fmt"

type file struct {
	name string
}

func NewFile(name string) iComponent {
	return &file{
		name: name,
	}
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}
