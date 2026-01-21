package errs

var _ error = String("")

// String is a typed string error value.
type String string

// Error returns the string contents.
func (e String) Error() string {
	return string(e)
}

// Is reports whether the target is also a String error, ignoring the value.
func (e String) Is(target error) bool {
	_, ok := target.(String)
	return ok
}
