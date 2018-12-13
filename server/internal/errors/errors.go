package errors

import e "errors"

var (
	// FatalError indicates a shutdown signal
	FatalError = e.New("FATAL_ERROR")
	// InternalServerError represents an internal server error
	InternalServerError = e.New("INTERNAL_SERVER_ERROR")
	// NotFoundError indicates that a resource was unable to be located
	NotFoundError = e.New("NOT_FOUND_ERROR")
	// UnauthorizedError represents a permission error
	UnauthorizedError = e.New("UNAUTHORIZED_ERROR")
)
