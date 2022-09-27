package main

import (
	"fmt"
)

type turnOff struct {
	device tver
}

func NewTurnOff(device tver) commander {
	return &turnOff{
		device: device,
	}
}

func (t *turnOff) execute() error {
	fmt.Println("turn off tv")
	return t.device.off()
}
