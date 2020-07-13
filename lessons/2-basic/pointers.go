package main

import (
	"fmt"
	"net/http"
)

func main() {
	client :=&http.Client{}
	resp,err :=client.Get("http://www.baidu.com")
	fmt.Println(resp, err)
}
