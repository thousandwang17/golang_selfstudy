package main

import "fmt"

type usb struct{}

type iusb interface {
	InsertIntoUSBPort()
	ReadData() string
}

func NewUSB() iusb {
	return &usb{}
}

func (u *usb) InsertIntoUSBPort() {
	fmt.Println("USB insertIntoComputerPorts")
}

func (l *usb) ReadData() string {
	return "i'm USB"
}
