package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func WordCount(s string) map[string] int {
	words :=strings.Fields(s)
	m:=make(map[string]int)
	for _,word :=range words{
		m[word]++
	}
	return m
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	var result = WordCount("test test tds tesds")
	fmt.Println(result)
	fmt.Println(m)
	m["test"]=Vertex{10.000,-90.000}
	fmt.Println(m)
	fmt.Println(math.Nextafter(2,4))
}