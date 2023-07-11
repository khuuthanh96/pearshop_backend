package errors

import (
	"errors"
	"strings"
)

// App constant errors
var (
	ErrRecordNotFound   = errors.New("record not found")
	ErrConflictResource = errors.New("conflict resource")
	ErrForbidden        = errors.New("forbidden")
	ErrLockedResource   = errors.New("locked resource")
	ErrContextCancelled = errors.New("context is canceled")
)

// TypeError define system error type
type TypeError string

var (
	TypeInternal           TypeError = "TYPE_INTERNAL"
	TypeServiceUnavailable TypeError = "TYPE_SERVICE_UNAVAILABLE"
	TypeUnauthorized       TypeError = "TYPE_UNAUTHORIZED"
	TypeForbidden          TypeError = "TYPE_FORBIDDEN"
	TypeInvalidArgument    TypeError = "TYPE_INVALID_ARGUMENT"
	TypeNotFound           TypeError = "TYPE_NOT_FOUND"
	TypeConflict           TypeError = "TYPE_CONFLICT"
)

// Code define system error code
type Code string

const (
	CodeInternal       Code = "CODE_INTERNAL"
	CodeUnauthorized   Code = "CODE_UNAUTHORIZED"
	CodeForbidden      Code = "CODE_FORBIDDEN"
	CodeInvalidPayload Code = "CODE_INVALID_PAYLOAD"

	// internal user error code
	CodeInternalUserUnauthorized Code = "CODE_INTERNAL_USER_UNAUTHORIZED"
	CodeInternalUserForbidden    Code = "CODE_INTERNAL_USER_FORBIDDEN"
	CodeInternalUserExisted      Code = "CODE_INTERNAL_USER_EXISTED"

	// product
	CodeProductNotFound  Code = "CODE_PRODUCT_NOT_FOUND"
	CodeProductIDInvalid Code = "CODE_PRODUCT_ID_INVALID"
)

// SystemError define system error
type SystemError interface {
	Type() TypeError
	Code() Code
	Message() string
	Param() interface{}
	StatusCode() int
	Error() string
}

// SystemErrors an array of system errors
type SystemErrors []SystemError

// Error implements error interface
func (errs SystemErrors) Error() string {
	errsString := make([]string, 0, len(errs))
	for _, err := range errs {
		errsString = append(errsString, err.Error())
	}

	return strings.Join(errsString, "\n")
}
