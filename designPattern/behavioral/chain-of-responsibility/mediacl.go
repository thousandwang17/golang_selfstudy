package main

import (
	"fmt"
)

type medical struct {
	next IDepartment
}

func NewMedical() IDepartment {
	return &medical{}
}

func (c *medical) execute(data *patient) error {
	fmt.Println("medical do somthing")

	if c.next != nil {
		return c.next.execute(data)
	}
	return nil
}

func (c *medical) setNext(i IDepartment) {
	c.next = i
}
