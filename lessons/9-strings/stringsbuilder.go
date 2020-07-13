package main
//https://studygolang.com/articles/12796
import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var b strings.Builder
	n := 0
	var wait sync.WaitGroup
	for n < 1000 {
		wait.Add(1)
		go func() {
			b.WriteString("1")
			n++
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(len(b.String()))
}