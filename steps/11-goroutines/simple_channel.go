package main

import (
	"fmt"
	"strconv"
)

func SimpleChannel() {
	ch := make(chan string, 100)
	fmt.Println("channel length is ",len(ch))
	fmt.Println("start pushing data")
	ch <- "foo"
	fmt.Println("new channel size is ",len(ch))
	go func(){
	for index := 0; index < 100; index++ {
			fmt.Println("Pushing data .....")
			ch <- strconv.Itoa(index)
		}
	}()
	fmt.Println("new channel size is ",len(ch))
	msg :=<-ch
	fmt.Println("latest msg:",msg)	
	fmt.Println("next",<-ch)
	fmt.Println(len(ch))
}

func main() {
	SimpleChannel()
}
