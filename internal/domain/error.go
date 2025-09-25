package domain

import (
	"errors"
	"fmt"
	"net/http"
)

type GPError struct {
	code   GPErrorCode
	origin error
}

func NewGPError(code GPErrorCode) *GPError {
	return &GPError{
		code:   code,
		origin: nil,
	}
}

func NewGPErrorWithError(code GPErrorCode, origin error) *GPError {
	return &GPError{
		code:   code,
		origin: origin,
	}
}

func (e GPError) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("Code: %d, Message: %s", e.code, e.code.Message())
	}
	return fmt.Sprintf("Code: %d, Error: %s", e.code, e.origin.Error())
}

func (e *GPError) Append(msg string) *GPError {
	if e.origin == nil {
		e.origin = errors.New(msg)
	}
	e.origin = fmt.Errorf("%s: %w", msg, e.origin)
	return e
}

func (e *GPError) HttpStatusCode() int {
	if int(e.code) >= int(ErrCodeInternalServer) {
		return http.StatusInternalServerError
	}
	return http.StatusBadRequest
}

func (e *GPError) Message() string {
	if e.HttpStatusCode() >= http.StatusInternalServerError {
		return "Internal Server Error"
	}

	if e.HttpStatusCode() >= http.StatusUnauthorized {
		return "Unauthorized"
	}

	return e.code.Message()
}

func (e *GPError) ErrorCode() GPErrorCode {
	return e.code
}
