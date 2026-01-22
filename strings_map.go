package errs

import (
	"maps"
	"slices"
)

// StringsMap is a map of string slices representing field-specific error messages.
type StringsMap map[string][]string

// Error returns a formatted string representation of the StringsMap.
func (m StringsMap) Error() string {
	return stringifyMap(map[string][]string(m))
}

// Is reports whether the target is also a StringsMap error, ignoring the value.
func (m StringsMap) Is(target error) bool {
	t, ok := target.(StringsMap)
	if !ok {
		return false
	}

	return maps.EqualFunc(m, t, func(v1 []string, v2 []string) bool {
		return slices.Equal(v1, v2)
	})
}
