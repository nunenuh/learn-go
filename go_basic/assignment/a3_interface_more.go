package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])
	// os.Args[1]
	fname := os.Args[1]
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
