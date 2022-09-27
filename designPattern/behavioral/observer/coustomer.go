package main

import "fmt"

type customer struct {
	id int
}

func NewCustomer(i int) observer {
	return &customer{
		id: i,
	}
}

func (c *customer) update(s string) {
	fmt.Printf("%v got %s \n", c.id, s)
}

func (c *customer) getId() int {
	return c.id
}

func (c *customer) setId(i int) {
	c.id = i
}
