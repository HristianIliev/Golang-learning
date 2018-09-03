package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}

func main() {
	b := 255
	a := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("value of b is", b)

	// Passing to a function
	c := 58
	fmt.Println("value of a before function call is", c)
	d := &c
	change(d)
	fmt.Println("value of a after function call is", c)
}
