package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Printf("tick")
		case <-boom:
			fmt.Printf("BOOM!")
			return
		default:
			fmt.Printf("           .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	today:=time.Now()
	var t int = today.Day()
	switch t {
	case 5,10,15:
		fmt.Println("clean it")
	case 25,26,27:
		fmt.Println("have a rest")
	case 31:
		fmt.Println("No information")
	}
}