package utils

import "time"

const DateFormat = "02-01-2006"

func ParseDateString(date string) (time.Time, error) {
	// Parse the date string
	parsedTime, err := time.Parse(DateFormat, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
