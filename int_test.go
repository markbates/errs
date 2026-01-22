package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Int(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		val  Int
		exp  string
	}{
		{
			name: "positive integer",
			val:  Int(42),
			exp:  "42 error",
		},
		{
			name: "zero",
			val:  Int(0),
			exp:  "0 error",
		},
		{
			name: "negative integer",
			val:  Int(-7),
			exp:  "-7 error",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			act := tc.val.Error()

			r.Equal(tc.exp, act)
		})
	}
}

func Test_Int_Is(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		err    error
		target error
		exp    bool
	}{
		{
			name:   "same type and value",
			err:    Int(5),
			target: Int(5),
			exp:    true,
		},
		{
			name:   "same type different value",
			err:    Int(5),
			target: Int(10),
			exp:    false,
		},
		{
			name:   "different type",
			err:    Int(5),
			target: String("some error"),
			exp:    false,
		},
		{
			name:   "unwrap same error",
			err:    fmt.Errorf("wrapping: %w", Int(5)),
			target: Int(5),
			exp:    true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			act := errors.Is(tc.err, tc.target)

			r.Equal(tc.exp, act)
		})
	}

}
