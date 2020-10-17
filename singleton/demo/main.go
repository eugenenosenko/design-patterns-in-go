package main

import (
	"fmt"

	"design-patterns-in-go/singleton"
)

func main() {
	db := singleton.GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)

	cities := []string{"Seoul", "Mexico City"}
	//tp := GetTotalPopulation(cities)
	tp := singleton.GetTotalPopulationEx(singleton.GetSingletonDatabase(), cities)

	ok := tp == (17500000 + 17400000) // testing on live data
	fmt.Println(ok)

	names := []string{"alpha", "gamma"} // expect 4
	tp = singleton.GetTotalPopulationEx(&singleton.DummyDatabase{}, names)
	ok = tp == 4
	fmt.Println(ok)
}
