package factory

const (
	Developer = iota
	Manager
)

type SomeEmployee struct {
	Name, Position string
	AnnualIncome   int
}

func NewSomeEmployee(role int) *SomeEmployee {
	switch role {
	case Developer:
		return &SomeEmployee{
			Name:         "",
			Position:     "developer",
			AnnualIncome: 60000,
		}
	case Manager:
		return &SomeEmployee{
			Name:         "",
			Position:     "manager",
			AnnualIncome: 12000,
		}
	default:
		panic("unsupported employee role")
	}
}
