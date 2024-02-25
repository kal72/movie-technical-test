package utils

import "time"

const DateFormat = "2006-01-02 15:04:05"

func ParseDateString(date string) (time.Time, error) {
	// Parse the date string
	parsedTime, err := time.Parse(DateFormat, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
