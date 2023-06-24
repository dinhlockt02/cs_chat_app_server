package authmodel

import "fmt"

const ForgetPasswordEmail = "Reset your password"

func ForgetPasswordEmailBody(link string) string {
	return fmt.Sprintf(
		`
	<p><a href="%s">Click here to reset your password</a></p>
`, link)
}
