package main

import "fmt"

type contactInfo struct { 
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
		// alex := person{firstName: "Alex", lastName: "Anderson"}
		alex := person{
			firstName: "Alex",
			lastName: "Anderson",
			contactInfo: contactInfo{
				email: "alex@gmail.com",
				zip: 30310020,
		}}
		// alexPointer := &alex
		alex.updateName("Alexa")
		alex.print()

		name  := "bill"
		fmt.Println(&name) // The & turns a value into a pointer

		garota := "Carrey"
		updateName(garota)
		fmt.Println(garota)
}

func (pointerToPerson *person) updateName(firstName string) { // *person indica que a func só pode ser chamada com um pointer para person
	(*pointerToPerson).firstName = firstName // The * turns a pointer into a value
}

func updateName(n string) {
	n = "Cloey"
}

func (p person) print() {
	fmt.Printf("%+v", p)
	// fmt.Println(p)
}