package main

import (
	"fmt"
)

type cashier struct {
	next IDepartment
}

func NewCashier() IDepartment {
	return &cashier{}
}

func (c *cashier) execute(data *patient) error {
	fmt.Println("Cashier do somthing")

	if c.next != nil {
		return c.next.execute(data)
	}
	return nil
}

func (c *cashier) setNext(i IDepartment) {
	c.next = i
}
