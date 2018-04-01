package main
//overview golang data types
//http://www.runoob.com/go/go-data-types.html

import (
	"math/cmplx"
	"fmt"
)

//global declaration
var (
	ToBe      bool       = false
	TrueStuff            = true
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(ToBe)
	fmt.Println(TrueStuff)
	fmt.Println(MaxInt)
	fmt.Println(z)
}
