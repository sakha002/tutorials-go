package main

import "fmt"

func main() {

	var myGuy person

	myGuy.firstName = "jack"
	myGuy.lastName = "Jo"

	myGuyPointer := &myGuy

	myGuyPointer.updateName()

	myGuy.print()

}

type person struct {
	firstName string
	lastName  string
}

func (personPointer *person) updateName() {
	(*personPointer).firstName = "Joe"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
