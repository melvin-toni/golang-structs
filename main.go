package main

import "fmt"

type person struct {
	firstName string
	lastName  string
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
	alex := person{firstName: "Alex", lastName: "Anderson"}
	/* Update the value of struct */
	alex.firstName = "Alexx"
	alex.lastName = "Sanderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
