package pkg

import (
	"fmt"
	"time"
)

func IsDateValid(date string) bool {
	_, err := time.Parse(time.DateTime, date)
	return err == nil
}

func ParseDate(date string) (*time.Time, error) {
	if len(date) < 10 {
		return nil, fmt.Errorf("invalid date value, it must be (yyyy-mm-dd)")
	}
	dateTime, err := time.Parse("2006-01-02", date[:10])
	if err != nil {
		return nil, err
	}
	return &dateTime, nil
}
