package main

import "fmt"

/**
Functions, signature, return values, named results
*/

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(100, 23))
	region, continent := GetLocation("Shanghai")
	fmt.Println(region, continent)
}

func GetLocation(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Shanghai", "Beijing":
		region, continent = "China", "Asia"
	case "New York", "Boston":
		region, continent = "USA", "America"
	default:
		region, continent = "Unkonw", "Unkown"
	}
	return region, continent
}
