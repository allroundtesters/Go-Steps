package main

import (
	"bufio"
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	fileName := "../tmp/test.txt"
	//
	content, err := ioutil.ReadFile(fileName)

	if err!=nil {
		log.Fatalln("Eror Reading file ",fileName)
	}

	fmt.Println(string(content))

	filePath := "../tmp/test.txt"
	f,_:=os.Open(filePath)
	defer f.Close()

	scanner :=bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err :=scanner.Err();err!=nil{
		log.Fatal(err)
	}

	f,err = os.Create("output.txt")
	if err!=nil{
		log.Fatalln("Error Creating file:",err)
	}
	defer f.Close()

	for _,str := range []string{"apple","banana","fruit"} {
		bytes,err1 := f.WriteString(str+"\n")
		if err1 != nil{
			log.Fatalln("Error Writing string:",err1)
		}
		fmt.Println("%d",bytes)
	}

}