package errs

import (
	"maps"
	"slices"
)

// ErrorsMap is a map of error slices representing field-specific error messages.
type ErrorsMap map[string][]error

// Error returns a formatted string representation of the ErrorsMap.
func (m ErrorsMap) Error() string {
	return stringifyMap(map[string][]error(m))
}

// Is reports whether the target is also a ErrorsMap error, ignoring the value.
func (m ErrorsMap) Is(target error) bool {
	t, ok := target.(ErrorsMap)
	if !ok {
		return false
	}

	return maps.EqualFunc(m, t, func(v1 []error, v2 []error) bool {
		return slices.Equal(v1, v2)
	})
}
