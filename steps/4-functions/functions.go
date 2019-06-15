package main

import (
	"fmt"
	"math"
)

// Basic function accept one parameter
// the format is like function_name(parameter paramterType)
func Echo(s string) {
	fmt.Println(s)
}

func Say(s string) string {
	return "Hello " + s
}

//function with return type,return single named value
//return variable name could be specify in function 
// := is for new variable = for initial ones
func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return phrase
}

// func Say3(s string)(phrase string,err error){
// 	phrase="Hello "+s
// 	return
// }
//multiple parameters and return multiple return values,tuple,also named return value
func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return q, r
}

func main() {
	Echo("Hello World")
	fmt.Println(Say("Man"))
	q, r := Divide2(11, 3)
	fmt.Printf("Quotient:%v,Remainder %v \n", q, r)

}
