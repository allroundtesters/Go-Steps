package main

import "fmt"

func sum(a []int,c chan int){
	sum :=0
	for _,v :=range a {
		sum +=v
	}
	c<-sum
}

func main() {

	c :=make(chan int,2) // a channel with size two
	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(len(c))
	fmt.Println(<-c)
	fmt.Println(len(c))
}
