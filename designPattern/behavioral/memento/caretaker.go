package main

type caretaker struct {
	monetors []*monetor
}

func newCaretaker() *caretaker {
	return &caretaker{make([]*monetor, 0)}
}

func (c *caretaker) addMontors(m *monetor) {
	c.monetors = append(c.monetors, m)
}

func (c *caretaker) getMontors(index uint) *monetor {
	if uint(len(c.monetors)) >= index+1 {
		return c.monetors[index]
	}
	return nil
}
