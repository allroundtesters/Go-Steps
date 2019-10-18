package main

import (
	"fmt"
	"unsafe"
)

var block="package"

const const_str  = "const string"

const (
	a_c = "abc"
	n_c= len(a_c)
	c_c = unsafe.Sizeof(a_c)
	a_iota =iota
	n_iota =iota
	v_iota =iota
)
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
	noVarMsg :="msg"
	fmt.Println(noVarMsg)
	fmt.Println(const_str)

	fmt.Println(n_c)
	fmt.Println(c_c)
	fmt.Println(a_iota)
	fmt.Println(n_iota)
}