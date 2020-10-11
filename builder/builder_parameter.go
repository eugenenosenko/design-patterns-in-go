package builder

import (
	"fmt"
)

type email struct {
	from, subject, to, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.from = from

	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body

	return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to

	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject

	return b
}

func sendEmailImpl(email *email) {
	fmt.Println(email)
}

func SendEmail(action func(*EmailBuilder)) {
	b := EmailBuilder{}
	action(&b)
	sendEmailImpl(&b.email)
}
