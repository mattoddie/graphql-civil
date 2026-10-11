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
	// ErrorDateTimeMustBeString is returned when unmarshalling a non-string value.
	ErrorDateTimeMustBeString = errors.New("datetime must be a string")
)

// MarshalCivilDateTime marshals a civil.DateTime as a GraphQL string in the
// form "2006-01-02T15:04:05", with nine-digit fractional seconds appended
// when non-zero.
func MarshalCivilDateTime(dt civil.DateTime) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, `"`+dt.String()+`"`)
	})
}

// UnmarshalCivilDateTime unmarshals a GraphQL string such as
// "2006-01-02T15:04:05" into a civil.DateTime. The string is parsed with
// [civil.ParseDateTime].
func UnmarshalCivilDateTime(v any) (civil.DateTime, error) {
	s, ok := v.(string)
	if !ok {
		return civil.DateTime{}, ErrorDateTimeMustBeString
	}

	dt, err := civil.ParseDateTime(s)
	if err != nil {
		return civil.DateTime{}, fmt.Errorf("civil.ParseDateTime: %w", err)
	}

	return dt, nil
}
