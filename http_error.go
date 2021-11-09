package utils

import (
	"fmt"
	"github.com/go-errors/errors"
)

// HttpError defines the generic structure of an HTTP error response.
type HttpError struct {
	Code    int
	Err     error
	Message string
}

// Error will return the string representation of the HttpError.Message and Err, omitting the former if it is not set.
func (err *HttpError) Error() string {
	if err.Message != "" {
		return fmt.Sprintf("%s: %s", err.Message, err.Err.Error())
	}
	return err.Err.Error()
}

// StackTrace will return a string representation of Err's stack trace
func (err *HttpError) StackTrace() string {
	return errors.Wrap(err.Err, 1).ErrorStack()
}

// ErrorAndStack will return both the output of Error and StackTrace in one call
func (err *HttpError) ErrorAndStack() (string, string) {
	return err.Error(), err.StackTrace()
}
