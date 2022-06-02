package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sum int
	sync.RWMutex
}

func (c *counter) add(wg *sync.WaitGroup) {
	c.Lock()
	defer c.Unlock()
	defer wg.Done()
	c.sum++
}

func (c *counter) get() int {
	c.RLock()
	defer c.RUnlock()
	return c.sum
}

func main() {
	var wg sync.WaitGroup
	cnt := counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go cnt.add(&wg)
	}
	wg.Wait()
	defer fmt.Println(cnt.get())
}
