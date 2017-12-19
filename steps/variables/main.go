package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'



	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)

	var t,p int=67,89
	var n string
	var m float64
	var mn bool
	
	t=2
	p=10
	n="this is"
	m=12332.00
	mn=true
	fmt.Println("%v %v \n",t,p)
	fmt.Println("%v \n",n)
	fmt.Println("%v \n",m)
	fmt.Println("%v \n",mn)

	var message = "Hello World!"
	var a1, b1, c1 = 1, 2, 3

	fmt.Println(message, a1, b1, c1)
}