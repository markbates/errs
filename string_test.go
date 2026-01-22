package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_String(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	err := String("some error")
	r.Equal("some error", err.Error())
}

func Test_String_Is(t *testing.T) {
	t.Parallel()

	const (
		err1 = String("some error")
		err2 = String("another error")
	)

	tcs := []struct {
		name   string
		err    error
		target error
		exp    bool
	}{
		{
			name:   "same error",
			err:    err1,
			target: err1,
			exp:    true,
		},
		{
			name:   "different type",
			err:    err1,
			target: errors.New("standard error"),
			exp:    false,
		},
		{
			name:   "different value",
			err:    err1,
			target: err2,
			exp:    false,
		},
		{
			name:   "different type 2",
			err:    err1,
			target: StringsMap{},
			exp:    false,
		},
		{
			name:   "unwrap same error",
			err:    fmt.Errorf("wrapping: %w", err1),
			target: err1,
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
