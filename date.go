package civil

import (
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/civil"
	"github.com/99designs/gqlgen/graphql"
)

// errors
var (
	ErrorDateMustBeString = errors.New("date must be a string")
)

// MarshalCivilDate marshals a civil date
func MarshalCivilDate(d civil.Date) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// nolint:errcheck,gosec
		io.WriteString(w, `"`+d.String()+`"`)
	})
}

// UnmarshalCivilDate unmarshals a civil date
func UnmarshalCivilDate(v interface{}) (civil.Date, error) {
	s, ok := v.(string)
	if !ok {
		return civil.Date{}, ErrorDateMustBeString
	}

	dt, err := civil.ParseDate(s)
	if err != nil {
		return civil.Date{}, fmt.Errorf("civil.ParseDate: %w", err)
	}

	return dt, nil
}
