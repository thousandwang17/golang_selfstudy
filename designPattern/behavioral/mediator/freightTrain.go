package main

import "fmt"

type freightTrain struct {
	meditorer meditorer
}

func NewFreightTrain(m meditorer) tainer {
	return &freightTrain{m}
}

func (p *freightTrain) arrived() {
	if p.meditorer.canArrived(p) {
		fmt.Println("freightTrain arrived station")
		return
	}

	fmt.Println("freightTrain  blocked, waiting for station ")
}

func (p *freightTrain) depart() {
	p.meditorer.notifyAboutDepature()
}

func (p *freightTrain) permitArrival() {
	p.arrived()
}
