package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jimmy",
		lastName:  "Neutron",
		contact: contactInfo{
			email:   "jimmy.neutron@email.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("First Name: %+v", jim.firstName)
	fmt.Printf("Last Name: %+v", jim.lastName)
	fmt.Printf("E-mail: %+v", jim.contact.email)
	fmt.Printf("Zipcode: %+v", jim.contact.zipCode)
}
