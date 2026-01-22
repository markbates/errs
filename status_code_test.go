package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StatusCode(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		val  StatusCode
		exp  string
	}{
		{
			name: "specific status code",
			val:  StatusCode(404),
			exp:  "status: 404",
		},
		{
			name: "zero status code defaults to 200",
			val:  StatusCode(0),
			exp:  "status: 200",
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

func Test_StatusCode_Is(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		err    error
		target error
		exp    bool
	}{
		{
			name:   "same type and value",
			err:    StatusCode(404),
			target: StatusCode(404),
			exp:    true,
		},
		{
			name:   "same type different value",
			err:    StatusCode(404),
			target: StatusCode(200),
			exp:    false,
		},
		{
			name:   "different type",
			err:    StatusCode(404),
			target: errors.New("standard error"),
			exp:    false,
		},
		{
			name:   "unwrap same type and value",
			err:    fmt.Errorf("wrapping: %w", StatusCode(404)),
			target: StatusCode(404),
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

func Test_StatusCode_MarshalJSON(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		val  StatusCode
		exp  string
	}{
		{
			name: "specific status code",
			val:  StatusCode(404),
			exp:  "404",
		},
		{
			name: "zero status code defaults to 200",
			val:  StatusCode(0),
			exp:  "200",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			actBytes, err := tc.val.MarshalJSON()
			r.NoError(err)

			act := string(actBytes)

			r.Equal(tc.exp, act)
		})
	}
}
