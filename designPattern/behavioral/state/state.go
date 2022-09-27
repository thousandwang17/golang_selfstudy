package main

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(amount int) error
	dispenseItem() error
}
