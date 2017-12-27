package main

import (
	"fmt"
)


func main(){
	alphas := []string{"abc", "def", "ghi"}

	for i:=0;i<len(alphas);i++{
		fmt.Println(alphas[i])
		fmt.Println("%d:%s \n",i,alphas[i])
	}

	for i,val:=range alphas{
		fmt.Println("%d:%s \n",i,val)	
	}

	for i := range alphas {
		fmt.Println(i)
	}

	for _,val :=range alphas {
		fmt.Println(val)
	}

	//like while
	x:=0
	for x<10{
		fmt.Println(x)
		x++
	}
}