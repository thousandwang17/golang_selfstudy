package main

func main() {
	// reciver
	tv := NewTv()

	// command
	turnOn := NewTurnOn(tv)
	turnOff := NewTurnOff(tv)

	// snder
	remoteController := NewRemoteController(turnOn)
	tvBar := NewTvBar(turnOff)

	// active
	tvBar.press()
	remoteController.press()

}
