package builder

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city

	return b
}

func (b *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	b.person.StreetAddress = street

	return b
}

func (b *PersonAddressBuilder) WithPostCode(postCode string) *PersonAddressBuilder {
	b.person.Postcode = postCode

	return b
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName

	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome

	return b
}

func (b *PersonJobBuilder) As(position string) *PersonJobBuilder {
	b.person.Position = position

	return b
}
