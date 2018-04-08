## Golang Control Flow
***tester's notes***

- if/else
- for-loop
- for-case


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