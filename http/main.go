package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, error := http.Get("http://www.google.com")
	if error != nil {
		fmt.Println("error:", error)
		os.Exit(1)
	}

	// byteSize := make([]byte, 99_999)
	// response.Body.Read(byteSize)
	// fmt.Println(string(byteSize))

	// io.Copy(os.Stdout, response.Body)

	lw := logWriter{}
	io.Copy(lw, response.Body)
}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("Number of bytes:", len(byteSlice))
	return len(byteSlice), nil
}
