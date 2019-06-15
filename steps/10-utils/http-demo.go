package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	res,err= http.Get("http://www.baidu.com")
	if err!=nil{
		log.Fatal(e)
	}
	page,_ := ioutil.ReadAll(res.Body)
	fmt.Println(page)
}