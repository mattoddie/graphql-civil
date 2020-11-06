package civil

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
)

func TestMarshalCivilDateTime(t *testing.T) {
	for _, tc := range []struct {
		datetime time.Time
		expected string
	}{
		{
			datetime: time.Date(1991, time.June, 1, 0, 0, 0, 0, time.UTC),
			expected: "1991-06-01T00:00:00",
		},
		{
			datetime: time.Date(2020, time.November, 6, 12, 0, 0, 0, time.UTC),
			expected: "2020-11-06T12:00:00",
		},
		{
			datetime: time.Time{},
			expected: "0001-01-01T00:00:00",
		},
	} {
		t.Run(tc.expected, func(t *testing.T) {
			d := civil.DateTimeOf(tc.datetime)
			b := bytes.Buffer{}
			MarshalCivilDateTime(d).MarshalGQL(&b)
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestUnmarshalCivilDateTime(t *testing.T) {
	for name, tc := range map[string]struct {
		datetime interface{}
		expected civil.DateTime
		err      error
	}{
		"1991-01-01": {
			datetime: "1991-01-01T00:00:00",
			expected: civil.DateTimeOf(time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC)),
		},
		"2020-11-06": {
			datetime: "2020-11-06T12:00:00",
			expected: civil.DateTimeOf(time.Date(2020, time.November, 6, 12, 0, 0, 0, time.UTC)),
		},
		"2020": {
			datetime: "2020",
			expected: civil.DateTime{},
			err: fmt.Errorf(
				"civil.ParseDateTime: %w",
				&time.ParseError{Layout: "2006-01-02t15:04:05.999999999", Value: "2020", LayoutElem: "-"},
			),
		},
		"numeric": {
			datetime: 2020,
			expected: civil.DateTime{},
			err:      ErrorDateTimeMustBeString,
		},
	} {
		t.Run(name, func(t *testing.T) {
			d, err := UnmarshalCivilDateTime(tc.datetime)
			assert.Equal(t, tc.expected, d)
			assert.Equal(t, tc.err, err)
		})
	}
}
