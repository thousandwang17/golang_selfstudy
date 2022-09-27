package main

type tvBar struct {
	command commander
}

func NewTvBar(c commander) *tvBar {
	return &tvBar{c}
}

func (b *tvBar) press() error {
	return b.command.execute()
}
