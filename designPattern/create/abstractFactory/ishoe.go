package main

type iShoe interface {
	setLogo(logo string)
	setSize(num int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(l string) {
	s.logo = l
}

func (s *shoe) setSize(num int) {
	s.size = num
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}
