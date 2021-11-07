package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	zipCode int
}

type personWithAddress struct {
	firstName string
	lastName  string
	contactInfo
}

func (p personWithAddress) print() {
	fmt.Printf("%v\n", p)
}

func (p *personWithAddress) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	alex := person{"Alex", "Anderson"}
	john := person{firstName: "John", lastName: "Lennon"}
	var bob person
	bob.firstName = "Bob"
	bob.lastName = "Marley"
	var empty person

	fmt.Println(alex, john, bob)
	fmt.Printf("%+v,%v\n", empty, empty)

	fmt.Println()
	jim := personWithAddress{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}
