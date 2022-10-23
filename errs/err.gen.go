// generate code - DO NOT EDIT

package errs

import "fmt"

type UserError struct {
	Name string
	Msg  string
}

func (e *UserError) Error() string {
	return e.Msg
}

func UserNotFound(user string) error {
	return &UserError{
		Name: "UserNotFound",
		Msg:  fmt.Sprintf("User %q not found.", user),
	}
}

func OrderNotFound(order string) error {
	return &UserError{
		Name: "OrderNotFound",
		Msg:  fmt.Sprintf("Order %q not found.", order),
	}
}

func UserNameTooShort(user string, minimum int) error {
	return &UserError{
		Name: "UserNameTooShort",
		Msg:  fmt.Sprintf("User %q name too short, minimum is %d.", user, minimum),
	}
}
