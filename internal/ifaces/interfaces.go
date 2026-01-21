package ifaces

// Isable marks errors that provide a custom errors.Is match.
type Isable interface {
	error
	Is(target error) bool
}

// Asable marks errors that provide a custom errors.As match.
type Asable interface {
	error
	As(target any) bool
}

// Unwrappable marks errors that provide a custom errors.Unwrap result.
type Unwrappable interface {
	error
	Unwrap() error
}

// IsAsable groups Isable and Asable on a single type.
type IsAsable interface {
	error
	Isable
	Asable
}

// StatusCoder marks errors that expose an HTTP status code.
type StatusCoder interface {
	error
	StatusCode() int
}
