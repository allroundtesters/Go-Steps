package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{}
	v2 = Vertex{1,100}
	v3=Vertex{X:1}
	vp = &Vertex{2,3}
)

func main() {
	fmt.Print(Vertex{1,2})
	var v = Vertex{1,4}
	fmt.Println(v.X,v.Y)
	p:=&v
	p.X=1e9
	fmt.Println(p.X)
	fmt.Println(v,&v)
	fmt.Println(vp)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}
