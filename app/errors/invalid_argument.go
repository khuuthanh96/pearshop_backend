package errors

import (
	"fmt"
	"net/http"
)

type invalidArgumentErr struct {
	code    Code
	message string
	param   interface{}
}

// NewInvalidArgumentErr inits a system invalid argument error
func NewInvalidArgumentErr(code Code, message string, param interface{}) SystemError {
	return &invalidArgumentErr{code, message, param}
}

// Type returns error type
func (e *invalidArgumentErr) Type() TypeError {
	return TypeInvalidArgument
}

// Code returns error code
func (e *invalidArgumentErr) Code() Code {
	return e.code
}

// Message returns error message
func (e *invalidArgumentErr) Message() string {
	return e.message
}

// Param returns error param value
func (e *invalidArgumentErr) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *invalidArgumentErr) StatusCode() int {
	return http.StatusBadRequest
}

// Error implements error interface
func (e *invalidArgumentErr) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
		TypeInvalidArgument, e.Code(), e.Message(), e.Param())
}
