package main

import "fmt"

func main() {
	fmt.Println("counting")
	//defer sequence
	defer fmt.Println("World")
	fmt.Println("Hello")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}