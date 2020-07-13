package main

import (
	"fmt"
	"time"
)

/**
bool
string
Numeric types:
uint either 32 or 64 bits
int same size as uint
uintptr an unsigned integer large enough to store the uninterpreted bits of
a pointer value
uint8 the set of all unsigned 8-bit integers (0 to 255)
uint16 the set of all unsigned 16-bit integers (0 to 65535)
uint32 the set of all unsigned 32-bit integers (0 to 4294967295)
uint64 the set of all unsigned 64-bit integers (0 to 18446744073709551615)
int8 the set of all signed 8-bit integers (-128 to 127)
int16 the set of all signed 16-bit integers (-32768 to 32767)
int32 the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64 the set of all signed 64-bit integers
(-9223372036854775808 to 9223372036854775807)
float32 the set of all IEEE-754 32-bit floating-point numbers
float64 the set of all IEEE-754 64-bit floating-point numbers
complex64 the set of all complex numbers with float32 real and imaginary parts
complex128 the set of all complex numbers with float64 real and imaginary parts
byte alias for uint8
rune alias for int32 (represents a Unicode code point)
*/

func TypeConvertion() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
}
/**
type assertion
 */
func timeMap(y interface{}) {
	//type assertion
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

type Stringer interface {
	String() string
}
type fakeString struct {
	content string
}
// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}
func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	} }

func main() {
	foo := map[string]interface{}{
		"Matt": 42, }
	timeMap(foo)
	fmt.Println(foo)
	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello, Gophers")
}