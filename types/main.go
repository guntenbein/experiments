package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// StackTraced is the interface for errors keeping the trace
type StackTraced interface {
	StackTrace() errors.StackTrace
}

// Coded is an interface of something with the code
type Coded interface {
	Code() string
}

// ErrorCoded implementation of error with code
type ErrorCoded string

// Code is the function returning the error code
func (ec ErrorCoded) Code() string {
	return string(ec)
}

// SRCTError - SingleReasonCodedTracedError, the error that has a code, traced and
// the error message is the 'Prefix : Reason'
type SRCTError struct {
	Prefix string
	Reason string
	ErrorCoded
	StackTraced
}

// Error returns the string representation of the error
func (e SRCTError) Error() string {
	if e.Prefix == "" {
		return e.Reason
	}
	return fmt.Sprintf("%s : %s", e.Prefix, e.Reason)
}

type IgnoreSentry struct {}

func (i IgnoreSentry) SendToSentry() bool {
	return false
}

type SentryAware interface {SendToSentry() bool}

// IncorrectAuthTokenError is the error for incorrectly formed auth token
type IncorrectAuthTokenError struct{SRCTError; IgnoreSentry}

// NewIncorrectAuthTokenError ...
func NewIncorrectAuthTokenError(reason string) IncorrectAuthTokenError {
	return IncorrectAuthTokenError{SRCTError: SRCTError{
		"",
		reason,
		"SS",
		errors.Wrap(fmt.Errorf("error"), "traced").(StackTraced),

	}}
}

func main() {
	var ie interface{}
	ie = NewIncorrectAuthTokenError("some")
	sentryAware, ok := ie.(SentryAware)
	if ok && !sentryAware.SendToSentry() {
		fmt.Println("yarr!")
	}
}
