package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println(os.Args[1]) //compile project into executable file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}



// use go build to create executable (main)