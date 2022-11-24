package main

import "fmt"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	// Validation here...
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	// Validation here...
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	// Validation here...
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	// Validation here...
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	fmt.Println("Sending some email", email)
}

// This functions accepts a function with a builder argument in order to build
// the email via the builder API
func SendEmail(action func(*EmailBuilder)) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func builderParameterExample() {
	SendEmail(
		func(b *EmailBuilder) {
			b.
				From("foo@example.com").
				To("bar@example.com").
				Subject("Some subject").
				Body("Some content")
		},
	)
}
