package main

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
