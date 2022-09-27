package main

import (
	"fmt"
)

type docter struct {
	next IDepartment
}

func NewDoctor() IDepartment {
	return &docter{}
}

func (c *docter) execute(data *patient) error {
	fmt.Println("docter do somthing")

	if c.next != nil {
		return c.next.execute(data)
	}
	return nil
}

func (c *docter) setNext(i IDepartment) {
	c.next = i
}
