package pkg

import "time"

func IsDateValid(date string) bool {
	_, err := time.Parse(time.DateTime, date)
	return err == nil
}

func ParseDate(date string) (*time.Time, error) {
	dateTime, err := time.Parse(time.DateTime, date)
	return &dateTime, err
}
