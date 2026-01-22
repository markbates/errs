package errs_test

import (
	"errors"
	"fmt"

	"github.com/markbates/errs"
)

func ExampleInt() {
	// Create an Int error with a specific value.
	err := errs.Int(42)
	fmt.Println(err.Error())

	// Create an Int error with the zero value.
	err = errs.Int(0)
	fmt.Println(err.Error())

	// Output:
	// 42 error
	// 0 error
}

func ExampleStatusCode() {
	// Create a StatusCode error with a specific code.
	err := errs.StatusCode(404)
	fmt.Println(err.StatusCode())
	fmt.Println(err.Error())

	// Create a StatusCode error with the zero value, which defaults to 200.
	err = errs.StatusCode(0)
	fmt.Println(err.StatusCode())
	fmt.Println(err.Error())

	// Output:
	// 404
	// status: 404
	// 200
	// status: 200
}

func ExampleString() {
	err1 := errs.String("an error occurred")
	fmt.Println(err1.Error())

	err2 := errs.String("boom")
	fmt.Println(err2.Error())

	err3 := fmt.Errorf("wrapping: %w", err1)
	fmt.Println(err3.Error())

	fmt.Println(errors.Is(err3, err1))

	// Output:
	// an error occurred
	// boom
	// wrapping: an error occurred
	// true
}
func ExampleStringsMap() {
	// Create a StringsMap error with field-specific errors.
	err := errs.StringsMap{
		"server":   {"error2", "error1"},
		"database": {"error3"},
	}

	// Error printing is sorted by field name,
	// and each error is printed on its own line.
	// The messages for each field are also sorted.
	fmt.Println(err.Error())

	// Output:
	// database:
	// 	- error3
	// server:
	// 	- error1
	// 	- error2
}

func ExampleErrorsMap() {
	// Create an ErrorsMap error with field-specific errors.
	err := errs.ErrorsMap{
		"server":   {errs.String("error2"), errs.String("error1")},
		"database": {errs.String("error3")},
	}

	// Error printing is sorted by field name,
	// and each error is printed on its own line.
	// The messages for each field are also sorted.
	fmt.Println(err.Error())

	// Output:
	// database:
	// 	- error3
	// server:
	// 	- error1
	// 	- error2
}
