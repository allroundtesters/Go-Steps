/**
 *
 * Problem 2
 *
 * Each new term in the Fibonacci sequence is generated by adding the previous two terms.
 * By starting with 1 and 2, the first 10 terms will be:

 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

 * Find the sum of all the even-valued terms in the sequence which do not exceed four million.
 **/

package main

import (
	"fmt"
)
const MAX=40000
func isEven(num int) bool{
	if(num%2==0){
		return true
	}
	return false
}

func findInFibonacii(){
	f1:=1
	f2:=1
	sum:=0

	for f2<MAX {
		if isEven(f2) {
			sum+=f2
		}
		f1,f2=f2,f1+f2
	}

	fmt.Println("Total %d",sum)
}


func main(){
	findInFibonacii()
}