package reflects

import "fmt"

type TestCase struct {}

func (tc TestCase)Method1(){
	fmt.Println("this is method1")
}

func (tc TestCase) Method2(){
	fmt.Println("this is method2")
}
