package main

import "fmt"

func main() {
	mySlice := []int{2, 3, 4, 5}
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:2])

	cities := make([]string, 3) //init length
	cities[0] = "c1"
	cities[1] = "c2"
	//cities[2]="c3"
	fmt.Println("length:", len(cities))
	cities = append(cities, "append1")
	fmt.Println(cities)
	fmt.Println("length:", len(cities))
	fmt.Println("the empty one:", cities[2])

}
