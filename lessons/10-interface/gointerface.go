package main

import (
	"fmt"
	"reflect"
)

//https://juejin.im/post/5da2e31ef265da5b5a721f9f

type MyStruct string

func (ms MyStruct) SayHelloTo(person string) string {
	for idx,run :=range ms {
		fmt.Printf("%d,%c \n",idx,run)
	}
	return "this is from interface"+person
}


type Eater interface {
	Eat()
}

type Runner interface {
	Run()
}

type Animal interface {
	Eater
	Runner
}

// 这里定义一个Dog的struct，并实现eat方法和run方法，这样就实现了动物的接口
type Dog struct {
	Name string
}

func (dog *Dog) Eat() {
	fmt.Printf("%s is eating.", dog.Name)
}

func (dog *Dog) Run() {
	fmt.Printf("%s is running.", dog.Name)
}
type Cat struct {
	Name string
}
func (cat *Cat) Eat() {
	fmt.Printf("%s is eating.", cat.Name)
}

func (cat *Cat) Run() {
	fmt.Printf("%s is running.", cat.Name)
}

func TypeJudge(animals ...interface{}) {
	for index, animal := range animals {
		switch animal.(type) {
		case *Dog:
			fmt.Printf("第%d个参数是Dog类型\n", index)
		case *Cat:
			fmt.Printf("第%d个参数是Cat类型\n", index)
		default:
			fmt.Println("不确定类型")
		}
	}
}

func main() {
	u :=MyStruct("this is myStruct")
	fmt.Println(u.SayHelloTo("man"))
	fmt.Println(reflect.TypeOf(u))
	//g := u.(*Greeting) //error, not an interface
	//fmt.Println(g)
	//g=&MyStruct{"this is"}  //can't be string
	//greeting :=g.(*Greeting)
	var animal1 Animal
	animal1 = &Dog{"doggy"}
	animal1.Eat()
	animal1.Run()
	dog := animal1.(*Dog)
	dog.Eat()
	dog.Run()
	if dog, ok := animal1.(*Dog); ok {
		fmt.Println("convert success")
		dog.Run()
	} else {
		fmt.Println("convert fail")
	}



}
