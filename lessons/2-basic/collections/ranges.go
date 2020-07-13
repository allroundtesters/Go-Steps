package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		if pow[i] >= 16 {
			break
		}
	}
	fmt.Println(pow)

	//for map
	cities := map[string]int{
		"New York": 8336697,
		"Los Angeles": 3857799,
		"Chicago": 2714856, }
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}
