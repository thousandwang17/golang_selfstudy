package main

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func newNoItemState(v *vendingMachine) state {
	return &noItemState{v}
}

func (h *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (h *noItemState) addItem(count int) error {
	h.vendingMachine.setState(h.vendingMachine.hasItem)
	return fmt.Errorf("Item out of stock")
}

func (h *noItemState) insertMoney(amount int) error {
	return fmt.Errorf("Item out of stock")
}

func (h *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
