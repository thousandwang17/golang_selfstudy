package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func newHasMoneyState(v *vendingMachine) state {
	return &hasMoneyState{v}
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *hasMoneyState) addItem(num int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *hasMoneyState) insertMoney(amount int) error {
	return fmt.Errorf("already pay")
}

func (h *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}
