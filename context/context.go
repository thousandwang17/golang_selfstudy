package main

import (
	"context"
	"log"
	"time"
)

func test(ctx context.Context, id int) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Printf("%v 456", id)
				return
			default:
				log.Printf("%v 123", id)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}(ctx)
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	test(ctx, 0)
	time.Sleep(time.Second)

	time.Sleep(time.Second)

	log.Print("=====")
	test(ctx, 1)

	ctx1, cancel1 := context.WithCancel(ctx)

	test(ctx1, 11)
	test(ctx1, 12)
	test(ctx1, 13)
	time.Sleep(time.Second)
	cancel1()
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	test(ctx, 0)
	time.Sleep(time.Second)

	time.Sleep(time.Second)

	log.Print("=====")
	test(ctx, 1)

	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()

	test(ctx1, 11)
	test(ctx1, 12)
	test(ctx1, 13)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
}

func main() {

	withCancel()

	log.Print("=== Timeout ===")
	withTimeout()
}
