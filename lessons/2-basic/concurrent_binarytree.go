package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func Walk(t *Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *Tree, ch chan int) {
	if t != nil {
		recWalk(t.Left, ch)
		ch <- t.Value
		recWalk(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			// not the same size
			return false
		case !ok1:
			// both channels are empty
			return true
		case x1 != x2:
			// elements are different
			return false
		default:
			// keep iterating
		}
	}
}

func main() {

	ch := make(chan int)
	go Walk(New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
}
