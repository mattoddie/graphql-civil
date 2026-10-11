package civil

import (
	"bytes"
	"testing"
)

// Each fuzz test checks that any string accepted by an unmarshaller
// marshals to a quoted string that unmarshals back to the same value.

func FuzzCivilDate(f *testing.F) {
	for _, s := range []string{"1991-06-01", "2020-11-06", "0001-01-01", "2020", "", "2020-13-45"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		d, err := UnmarshalCivilDate(s)
		if err != nil {
			return
		}
		var b bytes.Buffer
		MarshalCivilDate(d).MarshalGQL(&b)
		out := b.String()
		if len(out) < 2 || out[0] != '"' || out[len(out)-1] != '"' {
			t.Fatalf("MarshalCivilDate(%q) = %s, want quoted string", s, out)
		}
		got, err := UnmarshalCivilDate(out[1 : len(out)-1])
		if err != nil {
			t.Fatalf("UnmarshalCivilDate(%s): %v", out, err)
		}
		if got != d {
			t.Fatalf("round trip of %q: got %v, want %v", s, got, d)
		}
	})
}

func FuzzCivilDateTime(f *testing.F) {
	for _, s := range []string{"2020-11-06T12:30:00", "2020-11-06T12:30:00.5", "0001-01-01T00:00:00", "2020-11-06", ""} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		dt, err := UnmarshalCivilDateTime(s)
		if err != nil {
			return
		}
		var b bytes.Buffer
		MarshalCivilDateTime(dt).MarshalGQL(&b)
		out := b.String()
		if len(out) < 2 || out[0] != '"' || out[len(out)-1] != '"' {
			t.Fatalf("MarshalCivilDateTime(%q) = %s, want quoted string", s, out)
		}
		got, err := UnmarshalCivilDateTime(out[1 : len(out)-1])
		if err != nil {
			t.Fatalf("UnmarshalCivilDateTime(%s): %v", out, err)
		}
		if got != dt {
			t.Fatalf("round trip of %q: got %v, want %v", s, got, dt)
		}
	})
}

func FuzzCivilTime(f *testing.F) {
	for _, s := range []string{"12:30:00", "12:30:00.5", "00:00:00", "23:59:59.999999999", "24:00:00", ""} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		tm, err := UnmarshalCivilTime(s)
		if err != nil {
			return
		}
		var b bytes.Buffer
		MarshalCivilTime(tm).MarshalGQL(&b)
		out := b.String()
		if len(out) < 2 || out[0] != '"' || out[len(out)-1] != '"' {
			t.Fatalf("MarshalCivilTime(%q) = %s, want quoted string", s, out)
		}
		got, err := UnmarshalCivilTime(out[1 : len(out)-1])
		if err != nil {
			t.Fatalf("UnmarshalCivilTime(%s): %v", out, err)
		}
		if got != tm {
			t.Fatalf("round trip of %q: got %v, want %v", s, got, tm)
		}
	})
}
