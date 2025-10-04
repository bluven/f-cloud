package errorx

import (
	"net/http"
)

var (
	ErrNotAuthorized  = New(http.StatusUnauthorized, "You are not authorized")
	ErrForbidden      = New(http.StatusForbidden, "You are not authorized to access this resource")
	ErrRecordNotFound = New(http.StatusNotFound, "Record not found")
)

type LogicError struct {
	// http code. Also can be used for bussiness logic code
	Code    int
	Message string
}

func (e *LogicError) Error() string {
	return e.Message
}

func (e *LogicError) ToErrorResponse() ErrorResponse {
	return ErrorResponse{
		Error: e.Message,
	}
}

func New(code int, msg string) *LogicError {
	return &LogicError{
		Code:    code,
		Message: msg,
	}
}

func NewNotFound(msg string) error {
	return New(http.StatusNotFound, msg)
}

func NewConflict(msg string) error {
	return New(http.StatusConflict, msg)
}

func NewForbidden(msg string) error {
	return New(http.StatusForbidden, msg)
}

func NewAuthorized(msg string) error {
	return New(http.StatusUnauthorized, msg)
}

func NewBadRequest(msg string) error {
	return New(http.StatusBadRequest, msg)
}
