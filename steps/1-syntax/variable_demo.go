package main

import "fmt"

/**
Name Conventions:
const use upper case,start with underscore or letter
 */
const PRODUCT = "Product"
const PRICE = 500
const (
	name ="name"
	age =18
)
func main() {
	var i int = 4
	var s string = "s"
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)
	fmt.Println(age)
	fmt.Println(name)
}
