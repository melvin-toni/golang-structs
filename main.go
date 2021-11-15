package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo // another way of declaring embedded struct
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	/*
		Zero Value in Go
		In Go, if we did'nt assign any value, then there will be zero value which is the default value
		string = "", int = 0, float = 0, bool = false
		e.g.
		var alex person
		fmt.Println(alex)
		fmt.Printf("%+v", alex)
		if we print out code above then the result will be
		{  }
		{firstName: lastName:}%
	*/

	/* Initialize struct value */
	alex := person{
		firstName: "Alex",
		lastName:  "Sanderson",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "alex@anderson.com",
			zipCode: 43211,
		},
	}
	/*
		- Update the value of struct -
	*/
	// alex.firstName = "Alexx"
	// alex.lastName = "Sanderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	/*
		- Example of using receiver -
	*/
	// alex.print()

	/*
		- Exmaple of using pointer -
		Main rule of pointer
		Operator * => e.g. *pointerToPerson => turn MEM.ADDRESS into VALUE
		Operator & => e.g. &alex => turn VALUE into MEM.ADDRESS
	*/
	alexPointer := &alex
	alexPointer.updateName("Gary Roach")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
