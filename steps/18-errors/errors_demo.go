package main

import (
	"fmt"
	"math"
	"time"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}


func Sqrt(x float64)(float64,error) {
	if x<0 {
		return 0,ErrNegativeSqrt(x)
	}
	z :=1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z, nil
		} else {
			z = z - (z*z-x)/(z*2)
		}
	}
}

func run() error {
	return &MyError{
		What: "it doesn't work",
		When: time.Now(),
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Print(Sqrt(-3))
}
