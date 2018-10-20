package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//nesting contactInfo struct inside person struct
	contactInfo
}

func main() {
	//whenever we make structs we need to define all the properties that structs gonna have

	//creating a new person and passing the firstName and the lastName as strings
	//one thing to note here is that Go relies on the order of assignment
	//on line 16, the first arg passed here is assumed to be firstName
	//and the second args passed is assumed to be the lastName

	//alex := person{"Alex", "Anderson"}

	//to solve the issue or problems that we might face with ambiguity
	//we follow the following way to describe the structs

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex)

	//another way to declare structs
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	//another way to log structs
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 000000,
		},
	}
	jim.print()
}

//creating a receiver on p of type person
func (p person) print() {
	fmt.Printf("%+v", p)
}
