package main
import (
	"fmt"
)
func main(){
	num :=1
	if(num>3){
		fmt.Println("Many")
	}

	if num==1{
		fmt.Println("One")
	}else if num==2{
		fmt.Println("Two")
	}else {
		fmt.Println("More than Three")
	}

	switch {
	case num==1:
		fmt.Println("One")
	case num==2:
		fmt.Println("Two")
	case num>2:
	fmt.Println("More than two")
	default:
		fmt.Println("Don't know what happened") 
	}
}