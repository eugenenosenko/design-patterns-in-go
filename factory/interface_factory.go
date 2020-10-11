package factory

import (
	"fmt"
)

type person struct {
	name string
	age  int
}
type Person interface {
	SayHello()
}

// by returning an interface Person you're encapsulating struct person
// since you no longer can say p := NewPerson("Eugene", 28) and then update person's name.
func NewPerson(name string, age int) Person {
	return &person{name: name, age: age}
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s, I'm %d years old\n", p.name, p.age)
}
