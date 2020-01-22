package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", 1)
	}
}

func main(){
	f("direct")
	go f("goroutine")  //this is for goroutine concurrency
	go func(msg string){
		fmt.Println("In GoRoutines:",msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
