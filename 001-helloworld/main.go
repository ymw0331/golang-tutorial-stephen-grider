package main

//package == project == workspace

import "fmt" //standard lib package, shorten for format
//also alot of other packages, fmt, calculator, uploader (reusable package)

// func is short for function
func main() {
	fmt.Println("Testing 123")
}


// How code run in project
// Raw source code (main.go)

// RUN -> go run main.go -> Go Compiler -> Executable Program -> Runs the program

// BUILD -> go build main.go -> Go Compiler -> Executable Program
