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

	// this gets the address of Jim
	// this is the pointer for Jim
	jimPointer := &jim
	jimPointer.updateName("Jimboy")
	jim.print()
}

// pointerToPerson is the pointer of where person type was used
// *person is the value of the type person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerPerson is the value of the address of the variable that used person type
	// turns the address into a value
	(*pointerToPerson).firstName = newFirstName
}

// p is the receiver or reference of the type person
// any type person can use method print in dot syntax
func (p person) print() {
	fmt.Printf("First Name: %+v ", p.firstName)
	fmt.Printf("Last Name: %+v ", p.lastName)
	fmt.Printf("E-mail: %+v ", p.contactInfo.email)
	fmt.Printf("Zipcode: %+v ", p.contactInfo.zipCode)
}
