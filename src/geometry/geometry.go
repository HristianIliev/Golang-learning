package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var rectLen, rectWidth float64 = -6, 7

func init() {
	println("Main package initialised")
	if rectLen < 0 {
		log.Fatal("Length is less than zero")
	}

	if rectWidth < 0 {
		log.Fatal("Width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
