package main

import (
	"fmt"
	"sync"
)

// Create wait group for all goroutines
var wg sync.WaitGroup

// Concurency-safety type for write and read result
type resultLocker struct {
	result uint
	locker sync.RWMutex
}

// Method which adds square to result.
// Using in couple with WaitGroup for safety adding numbers
func (l *resultLocker) addSquare(wg *sync.WaitGroup, number uint) {
	l.locker.Lock()
	defer l.locker.Unlock()
	defer wg.Done()
	l.result += number * number
}

// Read result method
func (l *resultLocker) read() uint {
	l.locker.RLock()
	defer l.locker.RUnlock()
	return l.result
}

func main() {
	var result resultLocker
	arr := [5]uint{2, 4, 6, 8, 10}
	for _, elem := range arr {
		wg.Add(1)
		go result.addSquare(&wg, elem)
	}
	wg.Wait()
	fmt.Printf("%d\n", result.read())
}
