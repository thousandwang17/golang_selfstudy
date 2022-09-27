package main

import "fmt"

type ifatory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(s string) (ifatory, error) {
	switch s {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
