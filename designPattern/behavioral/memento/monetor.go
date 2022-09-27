package main

type monetor struct {
	status int
}

func newMonetor(status int) *monetor {
	return &monetor{status}
}

func (m *monetor) getSavedStatus() int {
	return m.status
}
