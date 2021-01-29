package common

import "net/http"

func NewError(message, code string, statusCode int) *Error {
	return &Error{message, code, statusCode}
}

type Error struct {
	message    string
	code       string
	statusCode int
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) StatusCode() int {
	if e.statusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.statusCode
}
