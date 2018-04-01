## Tester's Golang Notes -1 

- Golang Environment Setup
- Golang - Hello World

---

@title[Golang Environment Setup]
## Golang Environment Setup
- Install Golang
  * Go to golang.org
  * Download the latest go installation file
  * Setup GOROOT(How,just search it to solve)
- Setup GOPATH

```shell
WORKSPACE=`pwd`
echo "export GOPATH="${WORKSPACE}"">> ~/.zshrc
```

---

@title[Run Golang]
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

@title[Elements in Golang]
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
