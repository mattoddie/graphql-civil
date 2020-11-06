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
	ErrorTimeMustBeString = errors.New("time must be a string")
)

// MarshalCivilTime marshals a civil time
func MarshalCivilTime(t civil.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// nolint:errcheck,gosec
		io.WriteString(w, `"`+t.String()+`"`)
	})
}

// UnmarshalCivilTime unmarshals a civil time
func UnmarshalCivilTime(v interface{}) (civil.Time, error) {
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
