package main

import (
	"log"
	"net/rpc"
	"testing"
)

func BenchmarkRpc(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		client, _ := rpc.DialHTTP("tcp", "localhost:1234")
		t.ResetTimer()
		for pb.Next() {
			var result Result
			if err := client.Call("Cal.Square", 12, &result); err != nil {
				log.Fatal("Failed to call Cal.Square. ", err)
			}
		}
	})
}
