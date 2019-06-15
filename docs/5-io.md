## Golang IO

- read a file
- write to a file
- directory 

---

## Golang Read A File

leverage os,path/filepath standard lib:
Walk is like python's walk

```go

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "steps/" //run from the GOPATH

	filepath.Walk(path, Walker)
}

func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walk Error ......")
		return nil
	}
	if fi.IsDir() {
		fmt.Println("Directory:", fn)
	} else {
		fmt.Println("File: ", fn)
	}

	return nil
}

```

----

### Golang File Utils

- ioutils
- filepath
- bufio

