package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func BasicStructure_StackQueue() {
	//stack
	stack := make([]int, 0)
	stack = append(stack, 10)
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(v)
	fmt.Println(len(stack))

	//queue,FIFO
	queue := make([]int, 0)
	queue = append(queue, 10)
	v = queue[0]
	fmt.Println(v)
	queue = queue[1:]
	fmt.Println(len(queue))

	//
}

func BasicDict() {
	m := make(map[string]int)
	m["hello"] = 1
	delete(m, "hello")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func BasicSort(l []int) {
	sort.Ints(l)
	sort.Strings([]string{})
	s := []string{}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func BasicMath() {
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxFloat64)
}

func BasicConvertion() {
	s := "12345"
	num := int(s[0] - '0')
	str := string(s[0])
	b := byte(num + '0')
	fmt.Printf("%d%s%c\n", num, str, b)

	num,_=strconv.Atoi("s")
	str=strconv.Itoa(10)
	fmt.Println(num,str)
}
func main() {
	BasicMath()
	BasicConvertion()
}
