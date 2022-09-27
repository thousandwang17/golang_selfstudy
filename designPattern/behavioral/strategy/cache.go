package main

type node struct {
	pre, next *node
	val       string
}

type linkList struct {
	head, tail *node
}

type cache struct {
	hashmap      map[string]*node
	storage      linkList
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *cache {

	head := &node{}
	tail := &node{}

	head.next = tail
	tail.pre = head
	ls := linkList{
		head: head,
		tail: tail,
	}

	return &cache{
		storage:      ls,
		capacity:     0,
		maxCapacity:  1,
		evictionAlgo: e,
		hashmap:      make(map[string]*node),
	}
}

func (c *cache) setEvic(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) Exist(s string) bool {
	if _, exist := c.hashmap[s]; exist {
		return true
	}
	return false
}

func (c *cache) insert(s string) {

	if c.Exist(s) {
		return
	}

	if c.maxCapacity == c.capacity {
		c.evictionAlgo.evic(c)
		c.capacity--
	}

	node := &node{
		val: s,
	}

	// TODO : insert to link list
	c.hashmap[s] = node
	c.capacity++
}

func (c *cache) remove(s string) {
	// TODO : remove from link list
	delete(c.hashmap, s)
}
