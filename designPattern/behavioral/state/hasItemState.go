package main

import "fmt"

type hasItmeState struct {
	vendingMachine *vendingMachine
}

func newHasItemState(v *vendingMachine) state {
	return &hasItmeState{v}
}

func (h *hasItmeState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Printf("Item requestd\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequest)
	return nil
}

func (h *hasItmeState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItmeState) insertMoney(amount int) error {
	return fmt.Errorf("Please select item first")
}

func (h *hasItmeState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
