package main

type EvictionAlgo interface {
	evic(c *cache)
}
