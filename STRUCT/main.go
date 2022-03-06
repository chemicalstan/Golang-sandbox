package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

// of type person memoryaddress/pointer
func (personPointer *person) updateName(firstname string) {
	// person (&)memoryaddress (*)value
	*&personPointer.firstName = firstname
	// (*personPointer).firstName = firstname
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "Ekubo",
		contact: contactInfo{
			email:   "ebuka@gmail.com",
			zipCode: 100455,
		},
	}
	alex.updateName("Sam")
	fmt.Printf("%+v", alex)
}
