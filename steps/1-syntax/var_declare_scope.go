package main

import (
	"fmt"
)

var block="package"
func main(){
	fmt.Println("Current Block value is "+block)
	block:="function"
	fmt.Printf("this block is %s \n",block)
	{
		block:="inner"
		fmt.Printf("This block is %s.\n",block)
	}

	//init values
	var msg ="msg" 
	var a,b,c =1,2,3
	fmt.Printf("%s,%d,%d,%d",msg,a,b,c)

	var d,e,f = "s",false,32
	fmt.Println(d,e,f)

	//without var can be used in a func
	no_var_msg:="msg"
	fmt.Println(no_var_msg)
}