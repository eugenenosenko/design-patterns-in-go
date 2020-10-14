package main

import (
	"fmt"

	"design-patterns-in-go/prototype"
)

func main() {
	// scenario
	// create john
	john := &prototype.Person{
		Name:     "John",
		LastName: "Goodman",
		Address: &prototype.Address{
			StreetName: "London Street 12",
			City:       "London",
			Country:    "England",
		},
	}
	f := &prototype.Person{Name: "Friend", LastName: "Friendson", Address: &prototype.Address{
		StreetName: "Friendly St.",
		City:       "Friendburg",
		Country:    "Friendieland",
	}, Friends: []prototype.Person{}}

	// add a friend to john
	john.AddFriend(f)

	// now let's  create Jane who has the same last name
	// but by assigning john to jane, i'm copying also pointer to address struct
	jane := john
	jane.Name = "Jane"

	// we change jane's city and now john also has city as manchester
	jane.Address.City = "Manchester"

	// same address!
	fmt.Println(*jane.Address)
	fmt.Println(*john.Address)

	// let's create an unsafe deep copy
	doug := john.UnsafeDeepCopy()
	john.Address.City = "London"
	doug.Name = "Doug"
	doug.LastName = "Dougson"
	doug.Address.StreetName = "Calle del Rey 12"
	doug.Address.City = "Acapulco"
	doug.Address.Country = "Mexico"

	// changes are conserved
	fmt.Println(*doug, *doug.Address)
	fmt.Println(*john, *john.Address)

	// create safe deep copy via serialization
	mark := john.DeepCopy()
	mark.Name = "Mark"
	mark.LastName = "Caldo"
	fmt.Println(*mark)
	fmt.Println(*john)

	// prototype factory
	dorota := prototype.NewMainOfficeEmployee("Dorota", 100)
	magda := prototype.NewAuxOfficeEmployee("Magda", 200)

	fmt.Println(dorota)
	fmt.Println(magda)
}
