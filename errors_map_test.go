package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ErrorsMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name  string
		input ErrorsMap
		exp   string
	}{
		{
			name:  "nil map",
			input: nil,
			exp:   "",
		},
		{
			name:  "empty map",
			input: ErrorsMap{},
			exp:   "",
		},
		{
			name: "non-empty map",
			input: ErrorsMap{
				"field2": {String("error3")},
				"field1": {String("error2"), String("error1")},
			},
			exp: "field1:\n\t- error1\n\t- error2\nfield2:\n\t- error3",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			act := tc.input.Error()

			r.Equal(tc.exp, act)
		})
	}

}

func Test_ErrorsMap_Is(t *testing.T) {
	t.Parallel()

	m := ErrorsMap{
		"field": {String("error")},
	}

	tcs := []struct {
		name   string
		err    error
		target error
		exp    bool
	}{
		{
			name: "same map",
			err:  m,
			target: ErrorsMap{
				"field": {String("error")},
			},
			exp: true,
		},
		{
			name:   "different type",
			err:    m,
			target: String("some error"),
			exp:    false,
		},
		{
			name: "different value",
			err:  m,
			target: ErrorsMap{
				"field": {String("different error")},
			},
			exp: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			act := errors.Is(m, tc.target)
			r.Equal(tc.exp, act)

		})
	}
}

func Test_ErrorsMap_As(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name      string
		input     error
		expTarget ErrorsMap
		expOk     bool
	}{
		{
			name:      "correct type",
			input:     ErrorsMap{"field": {String("error")}},
			expTarget: ErrorsMap{"field": {String("error")}},
			expOk:     true,
		},
		{
			name:      "incorrect type",
			input:     String("some error"),
			expTarget: nil,
			expOk:     false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			var target ErrorsMap
			ok := errors.As(tc.input, &target)
			r.Equal(tc.expOk, ok)
			r.Equal(tc.expTarget, target)
		})
	}

}
