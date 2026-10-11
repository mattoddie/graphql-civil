package civil_test

import (
	"fmt"
	"os"

	civil "github.com/mattoddie/graphql-civil"
)

func ExampleMarshalCivilDate() {
	d, err := civil.UnmarshalCivilDate("2020-11-06")
	if err != nil {
		panic(err)
	}

	civil.MarshalCivilDate(d).MarshalGQL(os.Stdout)
	// Output: "2020-11-06"
}

func ExampleUnmarshalCivilDateTime() {
	dt, err := civil.UnmarshalCivilDateTime("2020-11-06T12:30:00")
	if err != nil {
		panic(err)
	}

	fmt.Println(dt.Date, dt.Time.Hour, dt.Time.Minute)
	// Output: 2020-11-06 12 30
}

func ExampleUnmarshalCivilTime() {
	_, err := civil.UnmarshalCivilTime(1230)
	fmt.Println(err)
	// Output: time must be a string
}
