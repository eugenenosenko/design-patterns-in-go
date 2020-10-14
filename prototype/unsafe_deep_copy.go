package prototype

type UnsafeDeepCopier interface {
	UnsafeDeepCopy() *Person
}

func (p *Person) UnsafeDeepCopy() *Person {
	c := *p
	c.Address = p.Address.UnsafeDeepCopy()
	copy(c.Friends, p.Friends)

	return &c
}
