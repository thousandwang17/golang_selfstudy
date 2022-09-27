package main

import "fmt"

type item struct {
	observerList map[int]observer
	name         string
	isStock      bool
}

type itemInterface interface {
	Register(o observer)
	Deregister(o observer)
	UpdateAvailability()
}

func NewItem(name string) itemInterface {
	return &item{
		name:         name,
		observerList: map[int]observer{},
	}
}

func (i *item) Register(o observer) {
	if _, exists := i.observerList[o.getId()]; !exists {
		i.observerList[o.getId()] = o
	}
}

func (i *item) Deregister(o observer) {
	if _, exists := i.observerList[o.getId()]; exists {
		delete(i.observerList, o.getId())
	}
}

func (i *item) notifyAll() {
	for k := range i.observerList {
		i.observerList[k].update(i.name)
	}
}

func (i *item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.isStock = true
	i.notifyAll()
}
