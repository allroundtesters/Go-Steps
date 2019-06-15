package main

import (
	"fmt"
)

// define dog type
type Dog struct {
	Name string
	Color string
}

func (d Dog) call(){
	fmt.Println("Here comes a %s dog,%s. \n",d.Color,d.Name)
}

func main(){
	Spot := Dog{Name:"Spot",Color:"brown"}

	var Rover Dog
	Rover.Name="rover"
	Rover.Color="loght blondish with some dark hair"

	Spot.call()
	Rover.call()
}