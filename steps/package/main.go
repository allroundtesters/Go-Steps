package main 

import (
	"fmt"
)
func main(){
	fmt.Println("This is in main package")
	fmt.Println(utils.Reverse("Hello World!"))
}
