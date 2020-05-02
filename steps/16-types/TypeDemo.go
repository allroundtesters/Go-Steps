package main

import (
	"fmt"
	"strings"
)

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

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}


func main() {
	fmt.Print(Vertex{1,2})
	var v = Vertex{1,4}
	fmt.Println(v.X,v.Y)
	p:=&v //reference
	p.X=1e9
	fmt.Println(p.X)
	fmt.Println(p,v,&v)
	fmt.Println(vp)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	game :=[][]string{
		[]string{"-", "-"},
		[]string{"-", "-"},
	}
	fmt.Println(game)
	fmt.Println(game[0][0])
	fmt.Println(game[0][1])
	printBoard(game)

}
