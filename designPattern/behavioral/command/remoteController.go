package main

type remoteController struct {
	command commander
}

func NewRemoteController(c commander) *remoteController {
	return &remoteController{c}
}

func (b *remoteController) press() error {
	return b.command.execute()
}
