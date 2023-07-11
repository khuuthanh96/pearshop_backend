package errors

import (
	"fmt"
	"net/http"
)

type notFoundErr struct {
	code    Code
	message string
	param   interface{}
}

func NewNotFoundErr(code Code, message string, param interface{}) SystemError {
	return &notFoundErr{code, message, param}
}

// Type returns error type
func (e *notFoundErr) Type() TypeError {
	return TypeNotFound
}

// Code returns error code
func (e *notFoundErr) Code() Code {
	return e.code
}

// Message returns error message
func (e *notFoundErr) Message() string {
	return e.message
}

// Param returns error param value
func (e *notFoundErr) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *notFoundErr) StatusCode() int {
	return http.StatusNotFound
}

// Error implements error interface
func (e *notFoundErr) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
		TypeNotFound, e.Code(), e.Message(), e.Param())
}
