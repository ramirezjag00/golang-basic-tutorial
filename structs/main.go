package main

import "fmt"

// struct == hash == object == dictionary
type person struct {
	firstName string
	lastName  string
}

func main() {
	// sample 1
	// person1 := person{ firstName: "Drey", lastName: "Dev" }
	// fmt.Println(person1.firstName) // "Drey"

	// sample 2
	// & yields pointer to the struct
	// samePerson := &person1
	// fmt.Println(samePerson.firstName) // "Drey"

	// sample 3
	var person1 person
	// string interpolation
	// print empty struct
	// {firstName: lastName:}%
	fmt.Printf("%+v", person1)

	person1.firstName = "Drey"
	person1.lastName = "Dev"
	fmt.Println(person1)
	// {Drey Dev}
}
