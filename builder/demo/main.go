package main

import (
	"fmt"

	"design-patterns-in-go/builder"
)

func main() {
	// regular builder
	html := builder.NewHTMLBuilder("ul").
		AddChild("li", "first child").
		AddChild("li", "second child").
		Build()
	fmt.Println(html)

	// faceted builder allows you to access use same builder structure to create complex objects
	const annualIncome = 120000
	person := builder.NewPersonBuilder().
		Lives().
		At("Grybowska 4").
		In("Warsaw").
		WithPostCode("00-891").
		Works().
		As("developer").
		At("OLX").
		Earning(annualIncome).
		Build()
	fmt.Println(person)
	// builder param allows you to avoid working with the object directly
	builder.SendEmail(func(b *builder.EmailBuilder) {
		b.Body("hello").
			From("Kennoby").
			To("Grievous").
			Subject("there")
	})
}
