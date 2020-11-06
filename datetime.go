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
	ErrorDateTimeMustBeString = errors.New("datetime must be a string")
)

// MarshalCivilDateTime marshals a civil datetime
func MarshalCivilDateTime(d civil.DateTime) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// nolint:errcheck,gosec
		io.WriteString(w, d.String())
	})
}

// UnmarshalCivilDateTime unmarshals a civil datetime
func UnmarshalCivilDateTime(v interface{}) (civil.DateTime, error) {
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
