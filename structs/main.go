package main

import "fmt"

type contactInfo struct { 
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
		// alex := person{firstName: "Alex", lastName: "Anderson"}
		alex := person{
			firstName: "Alex",
			lastName: "Anderson",
			contact: contactInfo{
				email: "alex@gmail.com",
				zip: 30310020,
		}}
		fmt.Println(alex)
		// fmt.Printf("%+v", alex)
}
