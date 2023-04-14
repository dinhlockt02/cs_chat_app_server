package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

type ValidationError []error

func (err ValidationError) Error() string {
	if len(err) == 0 {
		return ""
	}
	rs := "[\n"
	for i := range err {
		rs += err[i].Error() + "\n"
	}
	rs += "]"
	return rs
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (err *AppError) RootError() error {
	if e, ok := err.RootErr.(*AppError); ok {
		return e.RootErr
	}
	return err.RootErr
}

func (err *AppError) Error() string {
	return err.RootError().Error()
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal server error", err.Error(), "ErrInternal")
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewFullErrorResponse(http.StatusNotFound, err, fmt.Sprintf("%v not found", entity), err.Error(), "ErrNotfound")
}

var ErrInvalidObjectId = errors.New("invalid object id")
