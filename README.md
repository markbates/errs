# errs

Small, typed error values for status codes, ints, and strings, plus a tiny test helper package.

## Install

```bash
go get github.com/markbates/errs
```

## Types

### StatusCode

`StatusCode` carries an HTTP status code and formats it as `status: <code>`. A zero value
is treated as `500`.

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/markbates/errs"
)

func main() {
	err := errs.StatusCode(404)
	fmt.Println(err.StatusCode())                 // 404
	fmt.Println(err.Error())                      // status: 404
	fmt.Println(errors.Is(err, errs.StatusCode(0))) // true

	fmt.Println(errs.StatusCode(0).StatusCode()) // 500

	b, _ := json.Marshal(errs.StatusCode(0))
	fmt.Println(string(b)) // 500
}
```

### Int

`Int` is a typed integer error with the string form `<n> error`. `errors.Is` matches any
`Int` value by type.

```go
package main

import (
	"errors"
	"fmt"

	"github.com/markbates/errs"
)

func main() {
	err := errs.Int(418)
	fmt.Println(err.Error())                 // 418 error
	fmt.Println(errors.Is(err, errs.Int(0))) // true
}
```

### String

`String` is a typed string error. `errors.Is` matches any `String` value by type.

```go
package main

import (
	"errors"
	"fmt"

	"github.com/markbates/errs"
)

func main() {
	err := errs.String("boom")
	fmt.Println(err.Error())                      // boom
	fmt.Println(errors.Is(err, errs.String("")))  // true
	fmt.Println(errors.Is(err, errors.New("boom"))) // false
}
```

## Testing helpers

The `errstest` package provides a small assertion suite to check `Error()`, `errors.Is`,
`errors.As`, and `errors.Unwrap` behavior.

```go
import (
	"testing"

	"github.com/markbates/errs"
	"github.com/markbates/errs/errstest"
)

func TestString(t *testing.T) {
	errstest.AssertError(t, errstest.TestCase{
		Err: errs.String("nope"),
		Msg: "nope",
		IsTarget: func() error {
			return errs.String("")
		},
		AsTarget: func() any {
			return errs.String("")
		},
		UnwrapErr: nil,
	})
}
```

## License

See `LICENSE` if present in the repository.
