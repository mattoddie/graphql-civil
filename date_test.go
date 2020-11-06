package civil

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
)

func TestMarshalCivilDate(t *testing.T) {
	for _, tc := range []struct {
		date     time.Time
		expected string
	}{
		{
			date:     time.Date(1991, time.June, 1, 0, 0, 0, 0, time.UTC),
			expected: "1991-06-01",
		},
		{
			date:     time.Date(2020, time.November, 6, 12, 0, 0, 0, time.UTC),
			expected: "2020-11-06",
		},
		{
			date:     time.Time{},
			expected: "0001-01-01",
		},
	} {
		t.Run(tc.expected, func(t *testing.T) {
			d := civil.DateOf(tc.date)
			b := bytes.Buffer{}
			MarshalCivilDate(d).MarshalGQL(&b)
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestUnmarshalCivilDate(t *testing.T) {
	for name, tc := range map[string]struct {
		date     interface{}
		expected civil.Date
		err      error
	}{
		"1991-01-01": {
			date:     "1991-01-01",
			expected: civil.DateOf(time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC)),
		},
		"2020-11-06": {
			date:     "2020-11-06",
			expected: civil.DateOf(time.Date(2020, time.November, 6, 12, 0, 0, 0, time.UTC)),
		},
		"2020": {
			date:     "2020",
			expected: civil.Date{},
			err: fmt.Errorf(
				"civil.ParseDate: %w",
				&time.ParseError{Layout: "2006-01-02", Value: "2020", LayoutElem: "-"},
			),
		},
		"numeric": {
			date:     2020,
			expected: civil.Date{},
			err:      ErrorDateMustBeString,
		},
	} {
		t.Run(name, func(t *testing.T) {
			d, err := UnmarshalCivilDate(tc.date)
			assert.Equal(t, tc.expected, d)
			assert.Equal(t, tc.err, err)
		})
	}
}
