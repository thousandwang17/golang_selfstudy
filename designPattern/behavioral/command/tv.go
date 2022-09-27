package main

type tv struct {
	status bool
}

func NewTv() tver {
	return &tv{}
}

func (t *tv) on() error {
	t.status = true
	return nil
}

func (t *tv) off() error {
	t.status = false
	return nil
}
