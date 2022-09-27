package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

const RECVIVER = 10
const SEDDER = 10
const MAX_NUMBER = 100

type MaxHandel struct {
	Number int
	lock   sync.Mutex
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	rand.Seed(time.Now().UnixNano())

	// O2N()
	// N2O()

	test := make(chan int, 1)

	select {
	case test <- 1:
		// log.Println("1")
	default:
	}

	select {
	case test <- 2:
		// log.Println("2")
	default:
	}

	// log.Print("N2N start")

	N2N()
}

func O2N() {

	wg := sync.WaitGroup{}
	wg.Add(RECVIVER)

	timeout := time.After(time.Second)

	data := make(chan int, 100)
	go func() {
		for {
			select {
			case <-timeout:
				close(data)
				return
			default:
				<-time.After(10 * time.Millisecond)
				data <- rand.Intn(MAX_NUMBER)
			}
		}
	}()

	for i := 0; i < RECVIVER; i++ {
		go func(u int) {
			defer wg.Done()

			for d := range data {
				fmt.Print("%v", d)
				// log.Println(d)
			}

			// log.Printf("%v down", u)
		}(i)
	}

	wg.Wait()
}

func N2O() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	data := make(chan int, 100)
	stopch := make(chan struct{})

	//seder
	for i := 0; i < SEDDER; i++ {
		go func(u int) {
			for {
				value := rand.Intn(100)
				select {
				case <-stopch:
					//log.Printf("%v stop", u)
					return
				case data <- value:
					<-time.After(100 * time.Millisecond)
				}
			}
		}(i)
	}

	//recied
	go func() {
		defer wg.Done()

		max := 0

		for i := range data {
			if max == 100 {
				close(stopch)
				return
			}
			log.Println(i)
			max++
		}

		log.Println("end")
	}()
	wg.Wait()
}

func N2N() {
	wg := sync.WaitGroup{}
	wg.Add(RECVIVER)

	data := make(chan int, 100)
	stopch := make(chan struct{})
	toStop := make(chan string, 1)

	go func() {
		<-toStop
		//log.Printf("stop %v", stoper)
		close(stopch)
	}()

	// sender
	for i := 0; i < SEDDER; i++ {
		go func(u int) {
			for {
				value := rand.Intn(10000)

				if value == 0 {
					select {
					case toStop <- "sender":
						// log.Printf("SEDDER %d get 0", u)
					default:
					}
					return
				}

				//
				select {
				case <-stopch:
					// log.Printf("SEDDER stop1 %v", u)
					return
				default:
				}

				select {
				case <-stopch:
					// log.Printf("SEDDER stop2 %v", u)
					return
				case data <- value:
				default:
				}
			}

		}(i)
	}

	maxHandel := MaxHandel{}
	// reciver
	for i := 0; i < RECVIVER; i++ {
		go func(u int, handler *MaxHandel) {
			defer wg.Done()
			for {
				select {
				case <-stopch:
					// log.Printf("RECVIVER stop1 %v", u)
					return
				default:
				}

				select {
				case <-stopch:
					// log.Printf("RECVIVER stop2 %v", u)
					return

				case <-data:
					if handler.Number > 100 {
						select {
						case toStop <- fmt.Sprintf("receiver %v ", u):
							// log.Printf("receiver %v : Number over  100", u)
						default:
						}
						return
					}
					handler.lock.Lock()
					// log.Printf("%v", v)
					handler.Number++
					handler.lock.Unlock()
				default:
				}
			}
		}(i, &maxHandel)

	}

	wg.Wait()

	// log.Printf("end %v", maxHandel.Number)

}
