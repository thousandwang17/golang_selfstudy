package main

import "fmt"

func main() {

	// for
	words := []string{"Go", "语言", "高性能", "编程"}
	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// range
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}

	//
	persons := []struct{ no int }{{no: 1}, {no: 2}, {no: 3}}
	for _, s := range persons {
		s.no += 10
	}
	for i := 0; i < len(persons); i++ {
		persons[i].no += 100
	}
	fmt.Println(persons)
}
