package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Println("error", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

// go run printfile/main.go ./printfile/myfile.txt
