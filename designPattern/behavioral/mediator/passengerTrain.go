package main

import "fmt"

type passegerTrain struct {
	meditorer meditorer
}

func NewPassgerTrain(m meditorer) tainer {
	return &passegerTrain{m}
}

func (p *passegerTrain) arrived() {
	if p.meditorer.canArrived(p) {
		fmt.Println(" passegerTrain arrived station")
		return
	}

	fmt.Println("passegerTrain  blocked, waiting for station ")
}

func (p *passegerTrain) depart() {
	p.meditorer.notifyAboutDepature()
}

func (p *passegerTrain) permitArrival() {
	p.arrived()
}
