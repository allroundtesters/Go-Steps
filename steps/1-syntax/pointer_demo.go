package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var pointer_a int =0
	var ipAddress="10.0.0.1"

	var ip_pointer *string =&ipAddress
	var null_pointer *string

	fmt.Print(pointer_a)
	fmt.Println(&pointer_a)
	fmt.Println(&ipAddress)
	fmt.Println(ip_pointer)

	fmt.Println(null_pointer==nil)
	fmt.Println(null_pointer!=nil)
	fmt.Println(reflect.TypeOf(null_pointer))
}
