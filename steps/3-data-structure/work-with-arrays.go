package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v %p\n", s, len(x), cap(x), x,x)
}

var container=[]string{"zero","one","two"}
func main() {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)

	a=append(a,5)
	printSlice("a",a)

	container:=map[int]string{0:"zero",1:"one",2:"two"}
	fmt.Println(container[1])
}