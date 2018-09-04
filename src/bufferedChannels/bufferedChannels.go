package main

import (
	"fmt"
	"time"
)

// Second
func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to channel")
	}

	close(ch)
}

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Second
	channel := make(chan int, 2)
	go write(channel)
	time.Sleep(2 * time.Second)
	for v := range channel {
		fmt.Println("read value", v, "from channel")
		time.Sleep(2 * time.Second)
	}
}
