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

// Is returns true if the target is a StringsMap with the same keys and values.
func (m StringsMap) Is(target error) bool {
	t, ok := target.(StringsMap)
	if !ok {
		return false
	}

	return maps.EqualFunc(m, t, func(v1 []string, v2 []string) bool {
		return slices.Equal(v1, v2)
	})
}
