package main

import (
	"fmt"
)

func main() {
	var elements [] int
	var alphas = []string{"abc", "def", "ghi", "jkl"}
	elements = append(elements, 123)
	elements = append(elements, 456)
	fmt.Println("%v \n", elements)
	alphas = append(alphas, "pqr", "stu")
	fmt.Println("%v \n", alphas)
	fmt.Println("alphas size is :", len(alphas))
	fmt.Println(alphas[0])
	fmt.Println(alphas[1])

	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	if (elemExists("def", alphas)) {
		fmt.Println("Exists!!!")
	}

}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
