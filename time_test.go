package civil

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
)

func TestMarshalCivilTime(t *testing.T) {
	for _, tc := range []struct {
		date     time.Time
		expected string
	}{
		{
			date:     time.Date(1991, time.June, 1, 12, 0, 0, 0, time.UTC),
			expected: `"12:00:00"`,
		},
		{
			date:     time.Date(2020, time.November, 6, 12, 30, 10, 0, time.UTC),
			expected: `"12:30:10"`,
		},
		{
			date:     time.Time{},
			expected: `"00:00:00"`,
		},
	} {
		t.Run(tc.expected, func(t *testing.T) {
			d := civil.TimeOf(tc.date)
			b := bytes.Buffer{}
			MarshalCivilTime(d).MarshalGQL(&b)
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestUnmarshalCivilTime(t *testing.T) {
	for name, tc := range map[string]struct {
		time     interface{}
		expected civil.Time
		err      error
	}{
		"12:00:00": {
			time:     "12:00:00",
			expected: civil.TimeOf(time.Date(1991, time.January, 1, 12, 0, 0, 0, time.UTC)),
		},
		"12:30:00": {
			time:     "12:30:00",
			expected: civil.TimeOf(time.Date(2020, time.November, 6, 12, 30, 0, 0, time.UTC)),
		},
		"1230": {
			time:     "1230",
			expected: civil.Time{},
			err: fmt.Errorf(
				"civil.ParseTime: %w",
				&time.ParseError{Layout: "15:04:05.999999999", Value: "1230", LayoutElem: ":", ValueElem: "30"},
			),
		},
		"numeric": {
			time:     1230,
			expected: civil.Time{},
			err:      ErrorTimeMustBeString,
		},
	} {
		t.Run(name, func(t *testing.T) {
			d, err := UnmarshalCivilTime(tc.time)
			assert.Equal(t, tc.expected, d)
			assert.Equal(t, tc.err, err)
		})
	}
}
