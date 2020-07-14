package main

import "fmt"

/**
dict or hash in paython or java
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	celebs := map[string]int{
		"Nicolas Cage":       50,
		"Selena Gomez":       21,
		"Jude Law":           41,
		"Scarlett Johansson": 29}
	fmt.Printf("%#v", celebs)

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.000,-70.000}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}
