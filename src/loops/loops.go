package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}

		if i%2 == 0 {
			continue
		}

		fmt.Printf(" %d", i)
	}
}
