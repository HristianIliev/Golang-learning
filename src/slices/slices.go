package main

import "fmt"

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[:]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}

	fmt.Println("arrayafter", darr)

	i := make([]int, 5, 5)
	fmt.Println(i)
}
