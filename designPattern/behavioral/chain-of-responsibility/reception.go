package main

import (
	"fmt"
)

type reception struct {
	next IDepartment
}

func NewReception() IDepartment {
	return &reception{}
}

func (c *reception) execute(data *patient) error {
	fmt.Println("reception do somthing")

	if c.next != nil {
		return c.next.execute(data)
	}
	return nil
}

func (c *reception) setNext(i IDepartment) {
	c.next = i
}
