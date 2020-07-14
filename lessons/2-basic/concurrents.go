package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
https://talks.golang.org/2012/concurrency.slide#4
 */

//goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


//channels
//ch <- v // Send v to channel ch.
//v := <-ch // Receive from ch, and
//// assign value to v.

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v }
	c <- sum // send sum to c
}

func main() {
	go say("Hello")
	say("Hello-Sync")
	a := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int,10) //buffered channel
	go sum(a[:len(a)/2], ch)
	//fmt.Println(<-ch)
	go sum(a[len(a)/2:], ch)
	//fmt.Println(<-ch)
	x, y := <-ch, <-ch // receive data from channel
	fmt.Println(x,y,x+y)

	//buffered goroutines channel deadlock
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3
	c3 :=func() {c<-3}
	go c3()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	go c3()
	fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	/**
	The reason is that we are adding an extra value from inside a go routine, so
	our code doesn’t block the main thread. The goroutine is being called before
	the channel is being emptied, but that is fine, the goroutine will wait until the
	channel is available. We then read a first value from the channel, which frees a
	spot and our goroutine can push its value to the channel.
	 */
	fc := make(chan int, 10)
	go fibonacci(cap(fc), fc) //never block the main thread
	for i := range fc {
		fmt.Println(i)
	}
	fmt.Println("start select channel action")
	cq := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("start to send to channel")
			fmt.Println(<-cq)
		}
		quit <- 0
	}()
	fibonacciSelect(cq, quit)  //in main thread, how to know goroutines in main thread

	response := make(chan *http.Response, 1)
	errors := make(chan *error)
	go func() {
		resp, err := http.Get("http://matt.aimonetti.net/")
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()
	for {
		select {
		case r := <-response:
			fmt.Printf("%s", r.Body)
			return
		case err := <-errors:
			log.Fatal(*err)
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("Timed out!")
			return
		} }
}


func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ { c <- x
		x, y = y, x+y }
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	fmt.Println("Start F Channel")
	x, y := 0, 1
	for { //select different channel for event loop like
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x,y)
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default action when channel is ready!")
			time.Sleep(1 * time.Second)
		} } }