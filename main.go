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

	//& operator followed by a variable name gives the memory address of the value the variable is pointing at
	//below jimPointer will have the memory location of value stored for variable jim
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

//*pointer gives the value located at the memory address
//pointerToPerson gives the value at the  memory location of jimPointer

//notice the receiver type is changed to pointer type i.e *person
//so when this method is invoked the memory location of jimPointer is passed , and stored @ pointerToPerson whose type is *person
func (pointerToPerson *person) updateName(newFirstName string) {
	//below is an operator, it means we want to manipulate the value the pointer is referencing at
	(*pointerToPerson).firstName = newFirstName
}

//creating a receiver on p of type person
// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

//creating a receiver on pointer
func (p *person) print() {
	fmt.Printf("%+v", p)
}
