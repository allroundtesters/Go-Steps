package main

import (
	"fmt"
)

func print_number(){
	fmt.Println("This is for Number:")
	fmt.Println(42)
}

func print_binary(){
	fmt.Println("%d-%b\n",45,46)
}

func print_hex(){
	fmt.Println("%d \t %b \t %#x \n",42,42,42)
}
//tod: understand UTF-8/binary/hex

func main(){
	print_number()
	print_binary()
	print_hex()
}