package main

import "fmt"

type file struct {
	str string
}

type fileIterface interface {
	Inode
}

func NewFile(s string) fileIterface {
	return &file{
		str: s,
	}
}

func (f *file) Print(prefix string) {
	fmt.Println(prefix + f.str)
}

func (f *file) Clone() Inode {
	return NewFile(f.str + "_clone")
}
