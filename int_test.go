package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Int(t *testing.T) {
	t.Parallel()

	t.Run("Error()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e := Int(404)
		r.Equal("404 error", e.Error())
	})

	t.Run("Is()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e1 := Int(500)
		e2 := Int(404)
		e3 := fmt.Errorf("some other error")

		r.ErrorIs(e1, e2)
		r.True(errors.Is(e1, e2))
		r.True(e1.Is(e2))

		r.NotErrorIs(e1, e3)
		r.False(errors.Is(e1, e3))
		r.False(e1.Is(e3))
	})

	t.Run("As()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		var err error = Int(401)

		var target Int
		r.True(errors.As(err, &target))
		r.Equal(err, target)

		var target2 String
		r.False(errors.As(err, &target2))
	})

	t.Run("Unwrap()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e := Int(403)

		r.Nil(errors.Unwrap(e))
	})
}
