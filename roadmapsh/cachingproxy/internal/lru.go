package internal

import "fmt"

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*LRUNode[K, V]
	head     *LRUNode[K, V]
	tail     *LRUNode[K, V]
}

type LRUNode[K comparable, V any] struct {
	key   K
	value *V

	prev *LRUNode[K, V]
	next *LRUNode[K, V]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*LRUNode[K, V]),
	}
}

func (c *LRUCache[K, V]) Put(key K, value *V) {
	node := &LRUNode[K, V]{
		key:   key,
		value: value,
	}

	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
	if c.tail == nil {
		c.tail = node
	}

	c.cache[key] = node

	for len(c.cache) > c.capacity {
		c.evict()
	}
}

func (c *LRUCache[K, V]) evict() {
	if c.tail == c.head {
		fmt.Println("evict: ", c.tail.key)
		delete(c.cache, c.tail.key)
		c.head = nil
		c.tail = nil
		return
	}

	fmt.Println("evict: ", c.tail.key)
	delete(c.cache, c.tail.key)
	c.tail = c.tail.prev
	c.tail.next = nil

}

func (c *LRUCache[K, V]) Get(key K) (*V, bool) {
	v, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	if v.next != nil {
		v.next.prev = v.prev
	}

	if v.prev != nil {
		v.prev.next = v.next
	}

	v.prev = nil
	v.next = c.head

	c.head = v

	return v.value, true
}
