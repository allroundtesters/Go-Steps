package main

import "fmt"

type Employee interface {
	PrintName() string
	PrintAddress() string
	PrintSalary(basic int,tax int) float64
}

type Emp int
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id: \t",e)
	fmt.Println("Employee Name: \t",name)
}

func(e Emp) PrintSalary(basic int,tax int) float64{
	var salary = (basic*tax) /100
	return float64(basic - salary)
}

func main() {
	//var e1 Employee
	//e1 = Emp(1)
	//e1.PrintName("test")
}