package main

import "fmt"

func main() {
	caretaker := newCaretaker()

	origintor := NewOrigintor()

	origintor.setStatus(1)
	fmt.Printf("Originator Current State: %v\n", origintor.getStatus())
	caretaker.addMontors(origintor.createMontor())

	origintor.setStatus(2)
	fmt.Printf("Originator Current State: %v\n", origintor.getStatus())
	caretaker.addMontors(origintor.createMontor())

	origintor.setStatus(3)
	fmt.Printf("Originator Current State: %v\n", origintor.getStatus())
	caretaker.addMontors(origintor.createMontor())

	origintor.restoreFromMontor(caretaker.getMontors(1))
	fmt.Printf("Originator Current State: %v\n", origintor.getStatus())

	origintor.restoreFromMontor(caretaker.getMontors(0))
	fmt.Printf("Originator Current State: %v\n", origintor.getStatus())
}
