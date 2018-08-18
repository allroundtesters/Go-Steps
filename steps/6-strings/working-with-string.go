package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
)

func getTheFlag() *string{
	return flag.String("Hello","World!","again")
}
var name string
func init(){
	flag.CommandLine = flag.NewFlagSet("",flag.ExitOnError)
	flag.CommandLine.Usage=func(){
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name,"name","everyone","the greeting object.")	
}

func main() {
	
	str := "HI, This is String"
	lower := strings.ToLower(str)
	fmt.Println("This is lower case string:"+ lower)

	if strings.Contains(lower,"this") {
		fmt.Println("Yes, it is lower this")
	}

	fmt.Println("For 3-9 in Strings:"+str[3:9])
	fmt.Println("First 5:"+str[:5])
	sentence:="I am a sentence made up of words"
	words:=strings.Split(str," ")
	fmt.Println(words)

	fields:=strings.Fields(sentence)
	fmt.Println(len(fields))
	fmt.Println(fields)


	// var name = getTheFlag()
	flag.Parse()
	fmt.Printf(name)

}