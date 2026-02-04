package errs

// ErrorsMap is a map of error slices representing field-specific error messages.
type ErrorsMap map[string][]error

// Error returns a formatted string representation of the ErrorsMap.
func (m ErrorsMap) Error() string {
	return stringifyMap(map[string][]error(m))
}

// Is reports whether the target is also a ErrorsMap error, ignoring the value.
func (m ErrorsMap) Is(target error) bool {
	_, ok := target.(ErrorsMap)
	return ok
}
