package main

import (
	"fmt"
	"os"
)

func main() {
	doDefer()
}

func doDefer() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")
}

func openFileDefer() {
	file, err := os.Open("foo.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
}
