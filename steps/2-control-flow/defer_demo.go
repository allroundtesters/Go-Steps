package main

import "fmt"

func main() {
	defer fmt.Println("this is from defer function")
	fmt.Println("this is printing")

	for i:=0;i<10 ;i++  {
		defer fmt.Println("this is ",i)
	}

	fmt.Println("done")
}
