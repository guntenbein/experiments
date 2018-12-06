package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//////////////////////////////////////////////////////////////////////////////////

// IncorrectFieldTypeError is the error for missing request parameters
type IncorrectFieldTypeError string

// Error returns the string representation of the error
func (ift IncorrectFieldTypeError) Error() string {
	return fmt.Sprintf("incorrect field type: %s", string(ift))
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := errors.Wrap(IncorrectFieldTypeError("cause"),"wrapped error")

	if err, ok := err.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
	}
	if err2, ok := err.(IncorrectFieldTypeError); ok {
		fmt.Printf("%+v", err2)
	} else {
		fmt.Printf("NOOO")
	}
}