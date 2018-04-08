package main

//overview golang data types
//http://www.runoob.com/go/go-data-types.html

import (
	"math/cmplx"
	"fmt"
)

//global declaration
var (
	s         string     = "string"
	ToBe      bool       = false
	TrueStuff            = true
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
	a, b, c              = 1, 2, 3
)

var e, f = 123, 5567

const PI = 3.14

//b := "golang" //error here
//c := 4.17

func Declaration() {
	g, h := "123", 345
	fmt.Println(g, h)
}

func ConstDef() {
	const MAN = "MAN" //scope is in function
}
func main() {
	fmt.Println(ToBe)
	fmt.Println(TrueStuff)
	fmt.Println(MaxInt)
	fmt.Println(z, s, e, f)
	Declaration()
	fmt.Println(PI)
}
