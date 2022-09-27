package main

import "fmt"

type Account struct {
	name string
}

func NewAccount(name string) *Account {
	return &Account{
		name: name,
	}
}

func (a *Account) CheckAcount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
