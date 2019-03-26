package fasthttputil

import "errors"

var (
	// ErrInvalidBinding occurs when the binding destination is not a pointer to a struct
	ErrInvalidBinding = errors.New("Invalid binding")

	// ErrInvalidParameter occurs when a query parameter is provided but not of the correct type
	ErrInvalidParameter = errors.New("Query parameter provided but of incorrect type")
)
