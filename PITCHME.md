## Golang Basic Data Structure 

Golang Basic Data Structure:

- array
- dict
- structure

### Work With Array

- init
- append
- slice

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

*** when reslice, the pointer is different ***

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