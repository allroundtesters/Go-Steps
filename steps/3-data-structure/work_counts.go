package main

import (
	"fmt"
	"strings"
)

func CountWords(s string) map[string]int {
	fields := strings.Fields(s)
	numOfFields := len(fields)
	mapz := make(map[string]int)
	for y := 0; y < numOfFields; y++ {
		mapz[fields[y]]++
	}
	return mapz
}

func main() {
	result :=CountWords("this is thsi is test")
	fmt.Println(result)
}