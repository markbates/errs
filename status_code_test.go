package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StatusCode(t *testing.T) {
	t.Parallel()

	t.Run("Error()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e := StatusCode(404)
		r.Equal("status: 404", e.Error())
	})

	t.Run("Is()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e1 := StatusCode(500)
		e2 := StatusCode(404)
		var e3 error = fmt.Errorf("some other error")

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

		var err error = StatusCode(401)

		var target StatusCode
		r.True(errors.As(err, &target))
		r.Equal(err, target)

		var target2 String
		r.False(errors.As(err, &target2))
	})

	t.Run("Unwrap()", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		e := StatusCode(403)

		r.Nil(errors.Unwrap(e))
	})
	t.Run("status code zero defaults to 500", func(t *testing.T) {
		t.Parallel()
		r := require.New(t)

		err := StatusCode(0)
		r.Equal(500, err.StatusCode())
		r.Equal("status: 500", err.Error())
	})

	t.Run("json", func(t *testing.T) {
		t.Parallel()
		r := require.New(t)

		sc := StatusCode(0)
		r.Equal(500, sc.StatusCode())

		b, err := json.Marshal(sc)
		r.NoError(err)
		r.Equal("500", string(b))

		var decoded StatusCode
		err = json.Unmarshal([]byte("404"), &decoded)
		r.NoError(err)
		r.Equal(404, decoded.StatusCode())

	})
}
