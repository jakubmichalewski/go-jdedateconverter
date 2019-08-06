package jdedateconverter

import (
	"errors"
	"strconv"
	"time"
)

type JdeDate string

func (date JdeDate) ToDate() (*time.Time, error) {
	if len(date) != 6 {
		return nil, errors.New("incorrect date length")
	}
	d := string(date)
	_, err := strconv.Atoi(d)
	if err != nil {
		return nil, errors.New("incorrect date, can't convert to integer")
	}
	by, _ := strconv.Atoi(d[0:1])
	y, _ := strconv.Atoi(d[1:3])

	days, _ := strconv.Atoi(d[len(d)-3:])
	baseYear := 1900 + by*100
	year := baseYear + y

	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, days-1)
	return &t, nil
}
