package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		if pow[i] >= 16 {
			break
		}
	}
	fmt.Println(pow)

	//for map, k,v to loop
	cities := map[string]int{
		"New York": 8336697,
		"Los Angeles": 3857799,
		"Chicago": 2714856, }
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
	ReOrgNames()
}

func ReOrgNames(){
	var maxLength int
	for _,name:=range names{
		if l:=len(name);l>maxLength {
			maxLength =l
		}
	}
	output :=make([][]string,maxLength)
	for _,name :=range names {
		output[len(name)-1] = append(output[len(name)-1], name)
	}
	fmt.Printf("%v", output)}
