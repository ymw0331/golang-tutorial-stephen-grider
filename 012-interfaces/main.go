package main

// created empty structs so we can create function to work with it
type englishBot struct {
}

type spanishBot struct {
}

func main() {

}

func (eb englishBot) getGreeting() string {

	
}

// Interface based on custom types
// type englishBot struct     type spanishBot struct
// func (englishBot) getGreeting() string <- probably very diff logic in those functions -> func (spanishBot) getGreeting() string
// func printGreeting (eb englishBot)     <- these will porbably have identical logic ->	func printGreeting (sb spanishBot)
