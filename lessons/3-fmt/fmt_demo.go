package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

/**
%v is a generic placeholder. It will automatically convert your variable into a string with some default options. This is typically useful when printing primitives such as strings or numbers and you don’t need specific options.
%#v prints your variable in Go syntax. This means you could copy the output and paste it into your code and it’ll syntactically correct. I find this most useful when working with structs and slices because it will print out the types and field names.
%T prints your variable’s type. This is really useful for debugging if your data is passed as an interface{} and you want to see what its concrete type.
%d prints an integer in base-10. You can do the same with %v but this is more explicit.
%x and %X print an integer in base-16. One nice trick though is that you can pass in a byte slice and it’ll print each byte as a two-digit hex number.
%f prints a floating point number without an exponent. You can do the same with %v but this becomes more useful when we add width and precision flags.
%q prints a quoted string. This is useful when your data may have invisible characters (such as zero width space) because the quoted string will print them as escape sequences.
%p prints a pointer address of your variable. This one is really useful when you’re debugging code and you want to check if different pointer variables reference the same data.

https://medium.com/go-walkthrough/go-walkthrough-fmt-55a14bbbfc53
*/

type MyComponent struct {
	LogOutput io.Writer
}

var c MyComponent


func main() {
	var buf bytes.Buffer
	c.LogOutput = &buf
	if strings.Contains(buf.String(), "component finished") {
		_, _ = fmt.Fprintln("unexpected log output: %s", buf.String())
	}
}

