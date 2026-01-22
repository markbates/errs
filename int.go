package errs

import "fmt"

// Int is a typed integer error value.
type Int int

// Error formats the integer as "<n> error".
func (i Int) Error() string {
	return fmt.Sprintf("%d error", i)
}

// Is reports whether the target is also an Int error, ignoring the value.
func (i Int) Is(target error) bool {
	if t, ok := target.(Int); ok {
		return i == t
	}

	return false
}
