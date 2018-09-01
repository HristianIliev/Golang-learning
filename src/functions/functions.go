package main

import (
	"fmt"
)

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	areaForBlank, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f", areaForBlank)
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Can be rewritten as

func calculateBillRe(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Multiple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// Named return valuse
func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

// Blank identifier
func rectPropsBlank(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
