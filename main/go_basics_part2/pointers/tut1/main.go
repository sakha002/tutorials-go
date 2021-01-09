package main

import "fmt"

func main() {

	var myGuy person

	myGuy.firstName = "jack"
	myGuy.lastName = "Jo"

	myGuy.updateName()

	myGuy.print()

}

type person struct {
	firstName string
	lastName  string
}

func (p person) updateName() {
	p.firstName = "Joe"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
