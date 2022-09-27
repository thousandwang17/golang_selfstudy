package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	tx()
	lua()
}

func tx() {
	t1 := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			increment("tx_count")
		}()
	}

	wg.Wait()

	t2 := time.Now()
	fmt.Println(t2.Sub(t1), rdb.Get(ctx, "tx_count"))
}

const maxRetries = 10

// Increment transactionally increments the key using GET and SET commands.
func increment(key string) error {
	// Transactional function.
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// Actual operation (local in optimistic lock).

		// Operation is commited only if the watched keys remain unchanged.
		tx.Set(ctx, key, n+1, 0)
		return err
	}

	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		err := rdb.Watch(ctx, txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return err
	}

	return errors.New("increment reached maximum number of retries")

}

func lua() {
	t1 := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			keys := []string{"lua_count"}
			values := []interface{}{+1}
			incrBy.Run(ctx, rdb, keys, values...).Int()
		}()
	}

	wg.Wait()

	t2 := time.Now()
	fmt.Println(t2.Sub(t1), rdb.Get(ctx, "lua_count"))
}

var incrBy = redis.NewScript(`
local key = KEYS[1]
local change = ARGV[1]

local value = redis.call("GET", key)
if not value then
  value = 0
end

value = value + change
redis.call("SET", key, value)

return value
`)
