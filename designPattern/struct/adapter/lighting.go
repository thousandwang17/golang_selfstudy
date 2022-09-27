package main

import "fmt"

type lighting struct{}

type ILighting interface {
	InsertIntoLightningPort()
	ReadData() string
}

func NewLighting() ILighting {
	return &lighting{}
}

func (l *lighting) InsertIntoLightningPort() {
	fmt.Println("lighting insertIntoComputerPorts")
}

func (l *lighting) ReadData() string {
	return "i'm lighting"
}
