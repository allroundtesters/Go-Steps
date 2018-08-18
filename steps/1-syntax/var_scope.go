package main

import (
	"fmt"
	
)

var block="package"

func main(){
	block:="function"
	fmt.Println("this block is %s",block)
	{
		block:="inner"
		fmt.Printf("This block is %s.\n",block)
	}
}