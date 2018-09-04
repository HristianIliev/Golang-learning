package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello function")
	done <- true
}

// Second example
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10

	}

	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number & 10
		sum += digit * digit * digit
		number /= 10
	}

	cubeop <- sum
}

// Third example
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}

	close(chnl)
}

// Fourth example
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}

	close(dchnl)
}

func calcSquares4(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}

	squareop <- sum
}

func calcCubes4(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}

	cubeop <- sum
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function terminated")

	// Second example
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	// Third example
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received", v)
	}

	// Fourth example
	number4 := 589
	sqrch4 := make(chan int)
	cubech4 := make(chan int)
	go calcSquares4(number4, sqrch4)
	go calcCubes4(number4, cubech4)
	squares4, cubes4 := <-sqrch4, <-cubech4
	fmt.Println("Final output", squares4+cubes4)
}
