package decoder

import "strings"

// Error is a wrapper around a group of failed field errors
type Error struct {
	Errors map[string]error
}

// NewError returns a new Error with a valid map
func NewError() Error {
	return Error{
		Errors: map[string]error{},
	}
}

// Error returns a string of all the failed field names
func (e Error) Error() string {
	str := "Error decoding: "
	fields := []string{}
	for field, _ := range e.Errors {
		fields = append(fields, field)
	}

	return str + strings.Join(fields, ", ")
}
