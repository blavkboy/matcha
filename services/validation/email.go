package validation

import (
	"regexp"
	"fmt"
)

//ValidEmail will help us quickly evaluate whether the email that is sent to
//the server is a valid email address.
func ValidEmail(email string) bool {
	fmt.Println("Validating email")
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}
