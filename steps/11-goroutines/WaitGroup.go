package main

import (
	"sync"
	"fmt"
	"time"
)

func worker(id int,wg *sync.WaitGroup){
	defer wg.Done()
	
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main()  {
	var wg sync.WaitGroup
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go worker(i,&wg)	
	}
	wg.Wait()	
}