package factory

import (
	"fmt"
)

const _defaultEyeCount = 2

type Human struct {
	Name     string
	Age      int
	EyeCount int
}

// for simple structs curly braces might be fine,
// but sometimes you'd want to have a struct with default params
// or some internal constructor logic.
func NewHuman(name string, age int) *Human {
	return &Human{
		Name:     name,
		Age:      age,
		EyeCount: _defaultEyeCount,
	}
}

func (p *Human) String() string {
	return fmt.Sprintf("%s is %d old and has %d eyes", p.Name, p.Age, p.EyeCount)
}
