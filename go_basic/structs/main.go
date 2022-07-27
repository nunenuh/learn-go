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
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 83114,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	// name := "Bill"

	// fmt.Println(*&name)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

}

func (p *person) updateName(newFirstName string) {
	(p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("%+v", p)
}
