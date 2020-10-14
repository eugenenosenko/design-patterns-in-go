package prototype

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetName, City, Country string
}

type Person struct {
	Name, LastName string
	Address        *Address
	Friends        []Person
}

func (a *Address) UnsafeDeepCopy() *Address {
	return &Address{
		StreetName: a.StreetName,
		City:       a.City,
		Country:    a.Country,
	}
}

func (p *Person) AddFriend(f *Person) {
	p.Friends = append(p.Friends, *f)
}

func (p *Person) DeepCopy() *Person {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	_ = enc.Encode(p)

	dec := gob.NewDecoder(&buff)
	result := &Person{}
	_ = dec.Decode(result)

	return result
}

type DeepCopier interface {
	DeepCopy() *Person
}
