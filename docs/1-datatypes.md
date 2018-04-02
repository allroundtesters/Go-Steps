## Tester's Golang Notes-2

Golang Data Types:

- bool
- number
  * int
  * float32/float64
  * ......
- string

---

- inheritance
  * Pointer
  * Array
  * struct
  * Channel
  * Function
  * Slice
  * interface
  * Map
  * ........

---

- others
  * byte
  * rune, like int32
  * uint
  * int
  * uintptr

---

### Simple Example for Types


```go

import (
	"math/cmplx"
	"fmt"
)

//global declaration
var (
	ToBe      bool       = false
	TrueStuff            = true
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(ToBe)
	fmt.Println(TrueStuff)
	fmt.Println(MaxInt)
	fmt.Println(z)
}
```
----

## Golang Variable

Variable: 

```go

var identifier type

```
---

mulitple variations :

```go

var (
	s         string     = "string"
	ToBe      bool       = false
	TrueStuff            = true
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
	a, b, c              = 1, 2, 3
)

```

---

error:

```go

// only valid in a function
//b := "golang" //error here
//c := 4.17

```

### Golang CONST

const Usage:

```go

const PI = 3.14

```

