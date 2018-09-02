package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Recangle package initialised")
}

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
