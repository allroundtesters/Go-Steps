package main

import (
	"fmt"
	"math/cmplx"
)

func add(x int,y int) int  {
	return x+y
}

func swap(first string,second string) (string,string){
	return second,first
}

var c,python,java bool
var index int =0

var (
	ToBe bool = true
	MaxInt uint64 =1<<64-1
	z complex128=cmplx.Sqrt(-5+12)
)

const (
	BIG =1<<100
	Small = BIG >>99
)


func main() {
	var first_s,second_s string = "Hello","World"
	swap(second_s,first_s)
	fmt.Println(first_s,second_s)
	const Truth = true
	fmt.Println(Truth)
	var c,python,java = true,false,"No"
	fmt.Println(add(10,12))
	fmt.Println(c,python,java,index)

	fmt.Println(ToBe)
	fmt.Println(MaxInt)
	fmt.Println(z)
	z=100
	fmt.Println(z)

	fmt.Println(BIG)
	//BIG=0


}
