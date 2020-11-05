package civil

import (
	"fmt"
	"io"

	"cloud.google.com/go/civil"
	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
)

// MarshalCivilDate marshals a civil date
func MarshalCivilDate(d civil.Date) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// nolint:errcheck,gosec
		io.WriteString(w, d.String())
	})
}

// UnmarshalCivilDate unmarshals a civil date
func UnmarshalCivilDate(v interface{}) (civil.Date, error) {
	s, ok := v.(string)
	if !ok {
		return civil.Date{}, fmt.Errorf("Date must be a string")
	}

	dt, err := civil.ParseDate(s)
	if err != nil {
		return civil.Date{}, errors.Wrap(err, "civil.ParseDate")
	}

	return dt, nil
}
