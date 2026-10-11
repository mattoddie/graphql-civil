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
	// ErrorDateMustBeString is returned when unmarshalling a non-string value.
	ErrorDateMustBeString = errors.New("date must be a string")
)

// MarshalCivilDate marshals a civil.Date as a GraphQL string in the form
// "2006-01-02".
func MarshalCivilDate(d civil.Date) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, `"`+d.String()+`"`)
	})
}

// UnmarshalCivilDate unmarshals a GraphQL string such as "2006-01-02" into a
// civil.Date. The string is parsed with [civil.ParseDate].
func UnmarshalCivilDate(v any) (civil.Date, error) {
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
