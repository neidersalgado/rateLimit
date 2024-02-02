package error

import (
	"fmt"
)

type ErrorType uint

const (
	ValidationError ErrorType = iota + 1
	RateLimitError
	InvalidType
)

type CustomError struct {
	Type    ErrorType
	Message string
	Inner   error
}

func NewValidationError(msg string, inner error) *CustomError {
	return &CustomError{Type: ValidationError, Message: msg, Inner: inner}
}

func NewRateLimitError(msg string, inner error) *CustomError {
	return &CustomError{Type: RateLimitError, Message: msg, Inner: inner}
}

func NewInvalidRequest(msg string, inner error) *CustomError {
	return &CustomError{Type: InvalidType, Message: msg, Inner: inner}
}

func (e *CustomError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("[%d] %s: %s", int(e.Type), e.Message, e.Inner.Error())
	}
	return fmt.Sprintf("[%d] %s", int(e.Type), e.Message)
}

func (e *CustomError) Unwrap() error {
	return e.Inner
}
