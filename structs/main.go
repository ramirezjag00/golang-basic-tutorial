package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo // contactInfo contactInfo
}

func main() {
	jim := person{
		firstName: "Jimmy",
		lastName:  "Neutron",
		contactInfo: contactInfo{
			email:   "jimmy.neutron@email.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("First Name: %+v", jim.firstName)
	fmt.Printf("Last Name: %+v", jim.lastName)
	fmt.Printf("E-mail: %+v", jim.contactInfo.email)
	fmt.Printf("Zipcode: %+v", jim.contactInfo.zipCode)
}
