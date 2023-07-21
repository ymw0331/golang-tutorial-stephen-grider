// Struct is a data structure
// Equivalent as plain object in JS or POJO in Java

// Card Struct Field Definition
// suit -> <string> = "Spades"
// value -> <string> = "Ace"

// Types and its zero values
// string = ""
// int = 0
// float = 0
// bool = false

package main

import "fmt"

// define type of person (custom type)
type person struct {
	// 1. Tell go what fileds the person struct has
	firstName string
	lastName  string
	// contact   contactInfo // custom type
	contactInfo // same as above
}

// Type Person
// firstName -> <string>
// lastName -> <string>
// contact -> <contactInfo>

type contactInfo struct {
	email   string
	zipCode int
}

// Type ContactInfo (Embedding Struct)
// email -> <string>
// zip -> <int>

func main() {

	// 2. Create a new value of type person
	// alex := person{"Alex", "Anderson"} //not recommended
	// alex := person{firstName: "Alex", lastName: "Anderson"} //recommended
	// fmt.Println(alex.firstName)

	// another way of declaring struct
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) // Printf is for interpolation syntax

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	// fmt.Printf("%+v", jim)
	// fmt.Printf("\n")

	// &variable -> give me the memory address of the value this variable is pointing at
	// *pointer -> give me the value this memeory address is pointing at
	// 0001 | person{firstName:"Jim"...}
	// address | value
	// Turn address into value with *address
	// Turn value into address with &value

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")

	// jim.updateName("Jimmy")
	jim.print()

}

// value of receiver is a value of type *person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName  
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
