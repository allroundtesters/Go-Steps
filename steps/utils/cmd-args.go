package main

import (
	"flag"
	"fmt"
	"os"
)

var str string
var num int
var help bool

func main(){
	num_args:=len(os.Args)
	if num_args <2{
		fmt.Println(">> No args passed in")
	}

	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display Help")
	flag.Parse()

	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}
}
