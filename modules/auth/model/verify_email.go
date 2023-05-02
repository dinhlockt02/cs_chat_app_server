package authmodel

import "fmt"

const VerifyEmailTitle = "Verify your email"

func VerifyEmailBody(link string) string {
	return fmt.Sprintf(
		`
	<p><a href="%s">Click here to verify your email</a></p>
`, link)
}
