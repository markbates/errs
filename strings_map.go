package errs

// StringsMap is a map of string slices representing field-specific error messages.
type StringsMap map[string][]string

// Error returns a formatted string representation of the StringsMap.
func (m StringsMap) Error() string {
	return stringifyMap(map[string][]string(m))
}

// Is returns true if the target is a StringsMap.
func (m StringsMap) Is(target error) bool {
	_, ok := target.(StringsMap)
	return ok
}
