package main

func main() {

	cache := InitCache(&Lfu{})

	cache.insert("a")
	cache.insert("a")

	cache.insert("a")

	cache.setEvic(&Lru{})

	cache.insert("d")

	cache.setEvic(&Fifo{})

	cache.insert("e")
}
