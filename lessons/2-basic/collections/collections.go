package main

import "fmt"

func main() {
	var a[2] string
	a[0] ="Hello"
	a[1]="World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	
	b :=[2]string{"Hello","World"}
	fmt.Printf("%q",b)

	var ma[2][3]string
	for i := 0;i<2;i++{
		for j := 0; j<3;j++{
			ma[i][j]=fmt.Sprintf("a-%d%d",i,j)
		}
	}
	fmt.Printf("%q",ma)
}
