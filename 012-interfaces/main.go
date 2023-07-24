package main

import "fmt"

type bot interface {
	getGreeting() string
	// if there is any other type inside program, that has function getGreeting and string associated with it, that type is promoted as type bot
	
}

// implicit type, does not need to "implement"  -> type englishBot implement bot struct{}
// created empty structs, so we can create function to work with it
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// refactor version for printGreeting
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb spanishBot) { //go doesnt support function overloading
//	fmt.Println(sb.getGreeting())
//}

func (englishBot) getGreeting() string {
	//	very custom login for generating an english greeting
	return "Hi There!"
}

// func (sb spanishBot) getGreeting() string { //we can get rid of sb var and leave type only if not used
func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Interface based on custom type
// type englishBot struct     type spanishBot struct
// func (englishBot) getGreeting() string <- probably very diff logic in those functions -> func (spanishBot) getGreeting() string
// func printGreeting (eb englishBot)     <- these will probably have identical logic ->	func printGreeting (sb spanishBot)
