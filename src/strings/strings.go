package main

import (
	"fmt"
	"unicode/utf8"
)

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d \n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d \n", s, utf8.RuneCountInString(s))
}

func main() {
	name := "Hello World"
	printCharsAndBytes(name)
	fmt.Printf("\n")
	printCharsAndBytes(name)
	fmt.Printf("\n")
	printCharsAndBytes("Señor")

	// Constructing string from slice of runes
	byteSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(byteSlice)
	fmt.Println(str)

	// Length of string
	word1 := "Señor"
	length(word1)
}
