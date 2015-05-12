package decoder

import "strings"

type Error struct {
	Errors map[string]error
}

func NewError() Error {
	return Error{
		Errors: map[string]error{},
	}
}

func (e Error) Error() string {
	str := "Error decoding: "
	fields := []string{}
	for field, _ := range e.Errors {
		fields = append(fields, field)
	}

	return str + strings.Join(fields, ", ")
}
