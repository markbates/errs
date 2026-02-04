package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringsMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name  string
		input StringsMap
		exp   string
	}{
		{
			name:  "nil map",
			input: nil,
			exp:   "",
		},
		{
			name:  "empty map",
			input: StringsMap{},
			exp:   "",
		},
		{
			name: "non-empty map",
			input: StringsMap{
				"field2": {"error3"},
				"field1": {"error2", "error1"},
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

func Test_StringsMap_Is(t *testing.T) {
	t.Parallel()

	m := StringsMap{
		"field": {"error"},
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
			target: StringsMap{
				"field": {"error"},
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
			target: StringsMap{
				"field": {"different error"},
			},
			exp: true,
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

func Test_StringsMap_As(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name      string
		input     error
		expTarget StringsMap
		expOk     bool
	}{
		{
			name:      "correct type",
			input:     StringsMap{"field": {"error"}},
			expTarget: StringsMap{"field": {"error"}},
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

			var target StringsMap
			ok := errors.As(tc.input, &target)
			r.Equal(tc.expOk, ok)
			r.Equal(tc.expTarget, target)
		})
	}

}
