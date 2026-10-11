// Package civil provides gqlgen marshal and unmarshal functions for the
// civil.Date, civil.DateTime and civil.Time types from
// [cloud.google.com/go/civil].
//
// Map the GraphQL scalars to this package in gqlgen.yml:
//
//	models:
//	  Date:
//	    model: github.com/mattoddie/graphql-civil.CivilDate
//	  DateTime:
//	    model: github.com/mattoddie/graphql-civil.CivilDateTime
//	  Time:
//	    model: github.com/mattoddie/graphql-civil.CivilTime
//
// Values are exchanged as strings: dates as "2006-01-02", times as
// "15:04:05" and datetimes as "2006-01-02T15:04:05", with optional
// fractional seconds.
package civil
