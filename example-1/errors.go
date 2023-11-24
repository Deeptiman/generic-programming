package main

import (
	"errors"
)

type Error interface {
	error
	Code() int
	Message() string
}

type ErrorResponse struct {
	code    int
	message string
}

func (e ErrorResponse) Code() int {
	return e.code
}

func (e ErrorResponse) Message() string {
	return e.message
}

func (e ErrorResponse) Error() string {
	return e.message
}

func As(err error) Error {
	var e ErrorResponse
	if err != nil && errors.As(err, &e) {
		return e
	}

	return nil
}
