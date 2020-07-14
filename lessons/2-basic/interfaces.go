package main

import (
	"fmt"
	"os"
	"time"
)

type NameUser struct {
	FirstName, LastName string
}

func (u *NameUser) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work"}
}

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

func main() {
	u := &NameUser{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}
