package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) StringOf() string {
	return "dog name is " + d.Name
}

func NewDog(name string, age int) Dog {
	return Dog{
		Name: name,
		Age:  age,
	}
}

type Artist struct {
	Name, Genre string
	Songs int
}
func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	dog := NewDog("atest", 12)
	fmt.Println(dog)

	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	//tricky
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
