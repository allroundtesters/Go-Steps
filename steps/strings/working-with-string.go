package main

import (
	"fmt"
	"strings"
)

func main() {
	
	str := "HI, This is String"
	lower := strings.ToLower(str)
	fmt.Println("This is lower case string:"+ lower)

	if(strings.Contains(lower,"this")){
		fmt.Println("Yes, it is lower this")
	}

	fmt.Println("For 3-9 in Strings:"+str[3:9])
	fmt.Println("First 5:"+str[:5])
	sentence:="I am a sentence made up of words"
	words:=strings.Split(str," ")
	fmt.Println(words)

	fields:=strings.Fields(sentence)
	fmt.Println(fields)
}