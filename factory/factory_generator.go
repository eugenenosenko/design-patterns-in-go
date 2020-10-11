package factory

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach.
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

type Factory interface {
	Create(name string) *Employee
}

// structural approach
type EmployeeFactory struct {
	position     string
	annualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) Factory {
	return &EmployeeFactory{
		position:     position,
		annualIncome: annualIncome,
	}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}
