package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("C:/Users/Hristian/go/src/files/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Cotentents of file: ", string(data))

	// Line by line
	fptr := flag.String("fpath", "test.txt", "C:/Users/Hristian/go/src/files/test.txt")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal()
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
