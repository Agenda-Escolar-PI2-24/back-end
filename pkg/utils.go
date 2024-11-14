package pkg

import (
	"time"
)

func IsDateValid(date string) bool {
	_, err := time.Parse(time.DateTime, date)
	return err == nil
}

func ParseDate(date string) (*time.Time, error) {
    dateTime, err := time.Parse("2006-01-02", date)
    if err != nil {
        return nil, err
    }
    return &dateTime, nil
}