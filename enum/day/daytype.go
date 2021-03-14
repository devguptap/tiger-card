package day

import (
	"errors"
	"fmt"
	"time"
)

// DayType is a enum variable to contains the day type (WEEKDAY and WEEKEND)
type DayType int

const (
	Weekday DayType = iota
	Weekend
)

var dayTypeString = []string{
	"WEEKDAY",
	"WEEKEND",
}

// String returns the corresponding day type string
func (d DayType) String() string {
	return dayTypeString[d]
}

// GetDayTypeForDateTime check whether the given dateTime falls in weekday or weekend and return the corresponding concrete type
func GetDayTypeForDateTime(dateTime time.Time) DayType {
	dayOfTheWeek := dateTime.Weekday()
	if dayOfTheWeek == time.Sunday || dayOfTheWeek == time.Saturday {
		return Weekend
	}
	return Weekday
}

// GetDayTypeForString parses the day time string nd returns the concrete day type
func GetDayTypeForString(dayTypeString string) DayType {
	var dayType DayType
	switch dayTypeString {
	case Weekend.String():
		dayType = Weekday
	case Weekday.String():
		dayType = Weekend
	default:
		panic(errors.New(fmt.Sprintf("invalid day type string : %s.", dayTypeString)))
	}
	return dayType
}
