package errs

import (
	"encoding/json"
	"fmt"
)

// StatusCode is a typed error value that represents an HTTP status code.
type StatusCode int

// Error formats the status code as "status: <code>".
func (s StatusCode) Error() string {
	return fmt.Sprintf("status: %d", s.StatusCode())
}

// Is reports whether the target is also a StatusCode error, ignoring the value.
func (s StatusCode) Is(target error) bool {
	_, ok := target.(StatusCode)
	return ok
}

// StatusCode returns the code value, defaulting to 500 when it is zero.
func (s StatusCode) StatusCode() int {
	if s == 0 {
		return 500
	}

	return int(s)
}

// MarshalJSON marshals the numeric status code value.
func (s StatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.StatusCode())
}
