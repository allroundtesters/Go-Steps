package main

import (
	"fmt"
)

type PolicyType int32

const (
	PolicyMIN PolicyType = 0
	PolicyMAX PolicyType = 1
	PolicyMID PolicyType = 2
	PolicyAVG PolicyType = 3
)

func foo(p PolicyType){
	fmt.Printf("enum value: %v \n",p)
}

func (p PolicyType) String() string{
	switch (p) {
		case PolicyMIN: return "MIN"
		case PolicyMAX: return "MAX"
		case PolicyMID: return "MID"
		case PolicyAVG: return "AVG"
		default:         return "UNKNOWN"
		}
}

func main(){
	foo(PolicyAVG)
	fmt.Println(PolicyAVG)
}