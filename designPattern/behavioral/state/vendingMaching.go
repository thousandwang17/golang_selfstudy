package main

import "fmt"

type vendingMachine struct {
	hasItem     state
	itemRequest state
	noItem      state
	hasMoney    state

	currentState state
	itemCount    int
	itemPrice    int
}

func NewvendingMachine(itemPrice, itemCount int) state {

	res := &vendingMachine{}

	res.hasItem = newHasItemState(res)
	res.itemRequest = newItemRequestState(res)
	res.noItem = newNoItemState(res)
	res.hasMoney = newHasMoneyState(res)

	res.setState(res.hasItem)
	res.itemCount = itemCount
	res.itemPrice = itemPrice

	return res
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(amount int) error {
	return v.currentState.insertMoney(amount)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
