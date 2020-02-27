package main

import (
	"fmt"
)

func main(){
	c := make(chan int,2)
	c <- 1
	c <- 2
	t:= <-c
	fmt.Println("t:",t)
	fmt.Println(<-c)

	v,ok := <-c 
	print(v)
	print(ok)
	print("end ......")
}