package jdedateconverter

import (
	"testing"
	"time"
)

func TestJdeDate_ToDate(t *testing.T) {
	d := JdeDate("119023")
	tt := time.Date(2019, 1, 23, 0, 0, 0, 0, time.UTC)
	res, _ := d.ToDate()

	if *res != tt {
		t.Errorf("Got %v want %v", *res, tt)
	}
}

func TestJdeDate_ToDate_ReturnsError(t *testing.T) {
	d := JdeDate("119023a")
	_, e := d.ToDate()
	if e == nil {
		t.Error("Expected error got nothing")
	}
}

func TestJdeDate_ToDate_ReturnsError2(t *testing.T) {
	d := JdeDate("11902x")
	_, e := d.ToDate()
	if e == nil {
		t.Error("Expected error got nothing")
	}
}
