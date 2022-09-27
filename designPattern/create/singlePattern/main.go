package main

import "time"

func main() {

	for i := 0; i < 100; i++ {
		go getInstance2()
	}

	<-time.After(time.Second)
	getInstance()
}
