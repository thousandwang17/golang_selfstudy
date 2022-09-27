package main

import "fmt"

type itemRequestState struct {
	vendingMachine *vendingMachine
}

func newItemRequestState(v *vendingMachine) state {
	return &itemRequestState{v}
}

func (h *itemRequestState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (h *itemRequestState) addItem(num int) error {
	return fmt.Errorf("Item already requested")
}

func (h *itemRequestState) insertMoney(money int) error {
	if money < h.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", h.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	h.vendingMachine.setState(h.vendingMachine.hasMoney)
	return nil
}

func (h *itemRequestState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
