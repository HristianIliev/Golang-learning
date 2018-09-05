package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

// With channels
func incrementCh(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}

	w.Wait()
	fmt.Println("final value of x", x)

	// With channels
	x = 0
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCh(&wg, ch)
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}
