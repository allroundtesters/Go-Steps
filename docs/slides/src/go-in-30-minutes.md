## Learning Golang in 30 minutes

- Golang Environment Setup
- Golang Go mod 使用
- Golang - Hello World

---

## Golang Environment Setup
- Install Golang
  * Go to golang.org
  * Download the latest go installation file
  * Setup GOROOT(How,just search it to solve)
- Setup GOPATH
- Go mod管理Go项目的三方依赖

```shell
WORKSPACE=`pwd`
echo "export GOPATH=\"${WORKSPACE}\"">> ~/.zshrc
```

---

## Run Golang script

- helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

---

- go build

```sh
go build helloworld.go
```

- go run

```sh
go run helloworld.go 
```

---

## Elements in Golang

- package declaration
- imported package
- function
- statements
- comments

---

look at helloworld.go again:

```go
//package declaration
package main
//import package
import "fmt"
//function and main entry
func main() {
	//statement
	fmt.Println("Hello World!")
}
```

---

## Golang的类型

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

---

## Golang Variable 变量

golang变量定义: 

```go
//type不是强制声明
var identifier type

```

---

multiple variations :

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

---

### Golang 常量

const Usage:

```go
package main 
const PI = 3.14
```
---

## Golang Control Flow
***tester's notes***

- if/else
- for-loop
- for-case

---

### if else 

```go

num :=3
if num==1{
	fmt.Println("One")
}else if num==2{
    fmt.Println("Two")
}else {
	fmt.Println("More than Three")
}
	
```

### for-loop

for :

```go
	alphas := []string{"abc", "def", "ghi"}
	for i := 0; i < len(alphas); i++ {
		fmt.Println(alphas[i])
		fmt.Println("%d:%s \n", i, alphas[i])
	}
```
---

for-loop: index and val:

```go
	for i, val := range alphas {
		fmt.Printf("%d:%s \n", i, val)
	}
```

---

for-loop: index:

```go
	for i := range alphas {
		fmt.Println(i)
	}
```

---

for-loop: value only:

```go

	for _, val := range alphas {
		fmt.Println(val)
	}

```

---

for like while loop:

```go

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

```

---
## Golang Basic Data Structure 

Golang Basic Data Structure:

- array
- dict
- structure

---

### Work With Array

- init
- append
- slice

---

### Array: init

```go

var arrInt []int

```

---

### Array: append

here is an example of append usage:

```go
import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v %p\n", s, len(x), cap(x), x,x)
}

func main() {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)

	a=append(a,5)
	printSlice("a",a)
}
```
---

### cap VS len

- cap: how many elements of slice can contain
- len: how many elements of slice have now
- append return different slice when resliced

---

### work with slice

python like:

```go

var alphas = []string{"abc", "def", "ghi", "jkl"}
alpha2 := alphas[1:3]
fmt.Println(alpha2)	

```

---

for loop: 

```go

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

```

---
## Golang Function

- Define a Function
- Elements in a Function

### Define A Function

A function with:

1. ```func``` keyword to start a function block
2. define function Name, here is Say
3. define arguments: here is only one,s which is ```stirng``` type
4. define return type, here is string too

```go

func Say(s string) string{
	return "Hello "+ s
}

```

### Elements in a Function

- ```func``` keyword
- function name
- arguments
- return type, and multiple return values
- function code block
- could name a return name

```go

func Divide2(x,y float64) (q,r float64){
	q = math.Trunc(x/y)
	r = math.Mod(x,y)
	return q,r
}

```

---

with return name:

```go
func Say2(s string) (phrase string){
	phrase = "Hello "+ s
	return phrase
}
```
---
## Golang IO

- read a file
- write to a file
- directory 

---

## Golang Read A File

leverage os,path/filepath standard lib:
Walk is like python's walk

```go

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "steps/" //run from the GOPATH

	filepath.Walk(path, Walker)
}

func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walk Error ......")
		return nil
	}
	if fi.IsDir() {
		fmt.Println("Directory:", fn)
	} else {
		fmt.Println("File: ", fn)
	}

	return nil
}

```

----

### Golang File Utils

- ioutils
- filepath
- bufio

