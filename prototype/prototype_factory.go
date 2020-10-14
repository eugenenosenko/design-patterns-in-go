package prototype

import (
	"bytes"
	"encoding/gob"
)

type OfficeAddress struct {
	Suite            int
	StreetName, City string
}

type Employee struct {
	Name   string
	Office OfficeAddress
}

func (e *Employee) DeepCopy() *Employee {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(e)

	dec := gob.NewDecoder(buf)
	result := &Employee{}
	_ = dec.Decode(result)

	return result
}

// employee factory
// either a struct or some functions.
var mainOffice = Employee{
	"", OfficeAddress{0, "123 East Dr", "London"}}
var auxOffice = Employee{
	"", OfficeAddress{0, "66 West Dr", "London"}}

// utility method for configuring employee
//   ↓ lowercase
func newEmployee(proto *Employee,
	name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return result
}

//  factory functions
func NewMainOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

// prototype factory functions
func NewAuxOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}
