package main

import "fmt"

var (
	name string
	age int
	location string
)


//or variables

var n,l,a = "name","last",1

func main() {
	//var name,location ="name","location"
	name,location :="name","location"
	fmt.Println(name,location)

	var action =func(){
		fmt.Println("this is action method")
	}

	action()
}