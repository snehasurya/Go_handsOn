package main

import (
	"container/list"
	"fmt"
	"sync"
)

type LRUcache struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
	mu       sync.RWMutex
}
type Entry struct {
	key   string
	value any
}

func NewLRUcache(capacity int) *LRUcache {
	return &LRUcache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}
func (c *LRUcache) print() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println()
	for e := c.queue.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v --> ", e.Value)
	}
	fmt.Println("nil")
}
func (c *LRUcache) put(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if element, ok := c.items[key]; ok {
		c.queue.MoveToFront(element)
		element.Value.(*Entry).value = value
		return
	}
	if c.queue.Len() == c.capacity {
		oldestElement := c.queue.Back()
		if oldestElement != nil {
			c.queue.Remove(oldestElement)
			delete(c.items, oldestElement.Value.(*Entry).key)
		}
	}
	newEntry := &Entry{key: key, value: value}
	newElement := c.queue.PushFront(newEntry)
	c.items[key] = newElement
}

func (c *LRUcache) get(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	if element, ok := c.items[key]; ok {
		c.queue.MoveToFront(element)
		return element.Value.(*Entry).value
	}
	return nil
}

func (c *LRUcache) len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.queue.Len()
}
func main() {
	lru := NewLRUcache(3)
	lru.put("a", "sneha")
	lru.put("b", "pooja")
	lru.put("c", "gaurav")
	fmt.Printf("%+v  %+v  %+v", lru, lru.items, lru.queue)
	lru.print()
	lru.get("a")

	lru.put("d", "pankaj")

	lru.print()
	//time.Sleep(100 * time.Millisecond)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", n)
			lru.put(key, n)
			fmt.Printf("Put: %s\n", key)
		}(i)
	}
	//	runtime.NumGoroutine()
	wg.Wait()
	lru.print()
}
