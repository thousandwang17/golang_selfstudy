package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "accd , bacc"
	input, _ = rever(input)
	fmt.Println(input)
	input, _ = rever(input)
	fmt.Println(input)
}

func rever(s string) (string, error) {

	if !utf8.ValidString(s) {
		return s, errors.New("rever invaild utf8  ")
	}

	rn := []rune(s)
	l, r := 0, len(rn)-1
	for l < r {
		rn[l], rn[r] = rn[r], rn[l]
		l++
		r--
	}

	return string(rn), nil
}

// go test -fuzz=Fuzz -fuzztime 10s
// only work go 1.18 +
