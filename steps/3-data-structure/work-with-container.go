package main

import (
	"fmt"
)

var container1 = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2) {
		fmt.Printf("Error,unsupported container")
		return
	}

	fmt.Printf("The element is %q. (container type: %T)\n",
		container[1], container)

}
