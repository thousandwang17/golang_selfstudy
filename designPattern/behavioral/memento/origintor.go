package main

type origintor struct {
	status int
}

func NewOrigintor() *origintor {
	return &origintor{}
}

func (o *origintor) setStatus(s int) {
	o.status = s
}

func (o *origintor) getStatus() int {
	return o.status
}

func (o *origintor) createMontor() *monetor {
	return newMonetor(o.status)
}

func (o *origintor) restoreFromMontor(m *monetor) {
	if m == nil {
		return
	}

	o.status = m.getSavedStatus()
}
