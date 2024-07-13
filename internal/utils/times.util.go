package utils

import (
	"fmt"
	"time"
)

func getLocation() *time.Location {
	loc, error := time.LoadLocation("Asia/Bangkok")
	if error != nil {
		panic(error)
	}

	return loc
}

func TimeUTC(dateTime *string) time.Time {
	date := time.Now()

	if dateTime != nil {
		value, err := time.Parse(time.RFC3339Nano, *dateTime)
		if err != nil {
			fmt.Println(err)
		}

		date = value
	}

	return date.UTC()
}

func TimeLocal(dateTime *string) time.Time {
	loc := getLocation()

	date := time.Now()
	if dateTime != nil {
		value, err := time.Parse(time.RFC3339Nano, *dateTime)
		if err != nil {
			fmt.Println(err)
		}
		date = value
	}

	return date.In(loc)
}

func TimeLocalFormatISO8601(dateTime *time.Time) string {
	loc := getLocation()

	date := time.Now()
	if dateTime != nil {
		date = *dateTime
	}

	return date.In(loc).Format("2006-01-02T15:04:05+07:00")
}

func TimeUTCFormatISO8601(dateTime *time.Time) string {
	date := time.Now()
	if dateTime != nil {
		date = *dateTime
	}
	return date.Format("2006-01-02T15:04:05Z")
}
