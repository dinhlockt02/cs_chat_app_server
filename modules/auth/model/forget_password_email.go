package authmodel

import "fmt"

const ForgetPasswordEmail = "Reset your password"

func ForgetPasswordEmailBody(link string) string {
	return fmt.Sprintf(
		`
	<p>Our team has received an reset password request from you.</p>
	<p>Please click the link below or copy to your browser to reset your password.</p>
	<p>If you've not requested, please ignore this email.</p>
	<a href="%s">%s</a>
`, link, link)
}
