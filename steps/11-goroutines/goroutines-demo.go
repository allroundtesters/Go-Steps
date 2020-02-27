package main

import (
	"fmt"
	"time"
)

func say(s string){
	for i:=0;i<5 ;i++  {
		time.Sleep(100*time.Microsecond)
		fmt.Println(s)
	}
}

func PrintSlowly(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i, s)
		time.Sleep(300 * time.Microsecond)
	}
}

func main() {
	PrintSlowly("directly functioning", 3)
	go PrintSlowly("Red goroutine", 3)
	go PrintSlowly("Green goroutine", 3)

	// Call an anonymous function as a go routine.
	go func(ss string, nn int) {
		for i := 0; i < nn; i++ {
			fmt.Println(i, ss)
			time.Sleep(150 * time.Millisecond)
		}
	}("Yello goroutine", 3)

	var input string
	fmt.Scanln(&input) // Just push RETURN to finish the program.
	fmt.Println("DONE.")

	go say("world")
	say("hello")
}
