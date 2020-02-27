package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	m["test"]=Vertex{10.000,-90.000}
	fmt.Println(m)
	fmt.Println(math.Nextafter(2,4))
}