package main

import (
	"fmt"
	"reflect"
)

func ReadMethods(){

	tc :=reflect.TypeOf(&TestCase{})
	n :=tc.NumMethod()
	for i := 0; i < n;i++ {
		name := tc.Method(i).Name
		fmt.Printf(name)
	}
}

func main() {
	ReadMethods()
}
