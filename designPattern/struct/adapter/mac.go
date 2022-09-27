package main

import "fmt"

type mac struct{}

type imac interface {
	ReadPorts(l ILighting)
}

func NewMac() imac {
	return &mac{}
}

func (m *mac) ReadPorts(l ILighting) {
	fmt.Println(l.ReadData())
}
