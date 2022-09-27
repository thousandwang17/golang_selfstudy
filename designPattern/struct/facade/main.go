package main

import (
	"log"
)

func main() {
	walletFacade, err := NewWalletFacade("abc", 1111)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	if err := walletFacade.addMoneyToWallet("abc", 1111, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	if err = walletFacade.deductMoneyWallet("abc", 1111, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
