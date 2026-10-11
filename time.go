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
	// ErrorTimeMustBeString is returned when unmarshalling a non-string value.
	ErrorTimeMustBeString = errors.New("time must be a string")
)

// MarshalCivilTime marshals a civil.Time as a GraphQL string in the form
// "15:04:05", with nine-digit fractional seconds appended when non-zero.
func MarshalCivilTime(t civil.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, `"`+t.String()+`"`)
	})
}

// UnmarshalCivilTime unmarshals a GraphQL string such as "15:04:05" into a
// civil.Time. The string is parsed with [civil.ParseTime].
func UnmarshalCivilTime(v any) (civil.Time, error) {
	s, ok := v.(string)
	if !ok {
		return civil.Time{}, ErrorTimeMustBeString
	}

	dt, err := civil.ParseTime(s)
	if err != nil {
		return civil.Time{}, fmt.Errorf("civil.ParseTime: %w", err)
	}

	return dt, nil
}
