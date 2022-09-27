package main

import (
	"fmt"
)

type turnOn struct {
	device tver
}

func NewTurnOn(device tver) commander {
	return &turnOn{
		device: device,
	}
}

func (t *turnOn) execute() error {

	fmt.Println("turn on tv")
	return t.device.on()
}
