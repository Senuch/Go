package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	fileName := args[1]

	data, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	_, _ = io.Copy(os.Stdout, data)
}