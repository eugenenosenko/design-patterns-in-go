package main

import (
	"fmt"

	"design-patterns-in-go/factory"
)

func main() {
	// standard factory
	h := factory.NewHuman("Eugene", 28)
	fmt.Print(h)

	// interface factory
	p := factory.NewPerson("Eugene", 28)
	p.SayHello() // can't do p.name = "John"

	// factory generator

	// functional approach
	developerFactory := factory.NewEmployeeFactory("developer", 60000)
	managerFactory := factory.NewEmployeeFactory("manager", 120000)
	john := developerFactory("John Goodman")
	doug := developerFactory("Doug Warren")
	mike := developerFactory("Michael Corleone")
	chistian := managerFactory("Christian Hensen")
	fmt.Println(john, doug, mike, chistian)

	// structural approach.
	// similar to functional but you can change the positon and annual income at will
	bossFactory := factory.NewEmployeeFactory2("CEO", 222000)
	ceoEmployee := bossFactory.Create("Eugene Nosenko")
	fmt.Println(ceoEmployee)

	// prototype factory
	// creates a prototype for further changes
	someEmployee := factory.NewSomeEmployee(factory.Developer)
	someEmployee.Name = "Eugene Nosenko"
}
