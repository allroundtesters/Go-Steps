package main

import (
	"fmt"
	"runtime"
)
func main(){
	num :=1
	if(num>3){
		fmt.Println("Many")
	}

	if num==1{
		fmt.Println("One")
	}else if num==2{
		fmt.Println("Two")
	}else {
		fmt.Println("More than Three")
	}

	switch {
	case num==1:
		fmt.Println("One")
	case num==2:
		fmt.Println("Two")
	case num>2:
	fmt.Println("More than two")
	default:
		fmt.Println("Don't know what happened")
	}

	switch  os:=runtime.GOOS ; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}