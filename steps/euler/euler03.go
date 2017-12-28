/**
*
* Problem 3
*
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 
600851475143 ?
*
**/

package main

import (
	"fmt"
	"time"
)

var primes []int64

func main(){
	t: =time.Now()
	fmt.Println("Starting find largest prime factor")

	const num=600851475143
	buildPrimes(num)
	for index,value :=range primes {
		if num%value ==0 {
			fmt.Printf("Prime Factor: %d \n", prime)		}
	}

	fmt.Println("Total used timed: %d",time.Since(t))
}


func buildPrimes(num int64) {
	stop := int(math.Floor(math.Sqrt(float64(num))))
	primes = append(primes, 2) // seed primes
	for i := 3; i < stop; i += 2 {
		isPrime(int64(i)) // builds array
	}
}

func isPrime(num int64) bool {
	var stop = len(primes)
	for i:=0;i<stop;i+2{
		if(num%primes[i]==0){
			return false
		}

		if primes[i]*primes[i]>num{
			break
		}
	}

	primes=append(primes,num)
	return true
}