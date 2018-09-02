package main

import "fmt"

func main() {
	// var a [3]int
	// a[0] = 12
	// fmt.Println(a)

	// b := [...]float64{67.7, 89.8, 21, 78}
	// fmt.Println("length of b is", len(b))

	a := [...]float64{67.7, 89.9, 21, 78}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	fmt.Println("\nsum of all the elements of the array is", sum)

	// Multidimensional arrays
	a1 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printArray(a1)
	var b1 [3][2]string
	b1[0][0] = "apple"
	b1[0][1] = "samsung"
	b1[1][0] = "microsoft"
	b1[1][1] = "google"
	b1[2][0] = "AT&T"
	b1[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b1)

	// Slices
	c := [5]int{76, 77, 78, 79, 80}
	var d []int = c[1:4]
	fmt.Println(d)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
