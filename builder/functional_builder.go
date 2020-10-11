package builder

type Employee struct {
	name, position string
	salary         int64
}

type personMod func(employee *Employee)

type EmployeeBuilder struct {
	actions []personMod
}

func (b *EmployeeBuilder) Called(name string) *EmployeeBuilder {
	b.actions = append(b.actions, func(p *Employee) {
		p.name = name
	})

	return b
}

func (b *EmployeeBuilder) Earning(salary int64) *EmployeeBuilder {
	b.actions = append(b.actions, func(p *Employee) {
		p.salary = salary
	})

	return b
}
func (b *EmployeeBuilder) WorkingAs(position string) *EmployeeBuilder {
	b.actions = append(b.actions, func(p *Employee) {
		p.position = position
	})

	return b
}

func (b *EmployeeBuilder) Build() *Employee {
	e := Employee{}
	for _, a := range b.actions {
		a(&e)
	}

	return &e
}
