package main

import (
	"fmt"
	"math"
)

const s string="constant"

func main(){
	fmt.Println(s)
	const n = 500
	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(int64(n))
	fmt.Println(math.Sin(n))
}