package main
import "fmt"


const p = "death & taxes"

const (
	pi = 3.14
	language ="EN"
	a=iota
	b=iota
	c=iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10

)

func main(){
	const q=43
	fmt.Println("p-",p)	
	fmt.Println("q-",q)
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(GB)
	fmt.Println(KB)
	fmt.Println(TB)
}



