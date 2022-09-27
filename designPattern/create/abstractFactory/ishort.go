package main

type iShort interface {
	setLogo(logo string)
	setSize(num int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) setSize(num int) {
	s.size = num
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) getSize() int {
	return s.size
}
