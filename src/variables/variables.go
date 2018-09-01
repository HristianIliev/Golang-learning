package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		name   = "hristian"
		age    = 29
		height int
	)
	fmt.Println("my name is", name, "my age is ", age, "my height is", height)

	name1, age1 := "hristian", 29
	fmt.Println("my name is", name1, "my age is ", age1)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Min value is: ", c)
}
