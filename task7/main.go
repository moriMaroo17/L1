package main

import (
	"fmt"
	"sync"
)

// Concurency-safety map
type safeMap struct {
	data map[string]string
	sync.RWMutex
}

// Put key-value method
func (m *safeMap) put(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

// Get value by key method
func (m *safeMap) get(key string) string {
	m.RLock()
	defer m.RUnlock()
	return m.data[key]
}

// Starts 10 goroutines, which puts key-value into the map
func main() {
	var wg sync.WaitGroup
	smap := safeMap{data: make(map[string]string)}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			smap.put(fmt.Sprint(num), fmt.Sprint(num))
			wg.Done()
		}(i)
	}
	wg.Wait()
	// Print all key-values in the map
	fmt.Printf("%v\n", smap.data)
}
