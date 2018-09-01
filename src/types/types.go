package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("type of b is %T, size of b is %d", b, unsafe.Sizeof(b))

	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println("Sum:", sum)
}
