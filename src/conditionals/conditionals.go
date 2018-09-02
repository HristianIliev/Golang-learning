package main

import (
	"fmt"
)

func main() {
	if num := 10; num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
