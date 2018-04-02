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