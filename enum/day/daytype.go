package day

import (
	"errors"
	"fmt"
	"time"
)

type DayType int

const (
	Weekday DayType = iota
	Weekend
)

var dayTypeString = []string{
	"WEEKDAY",
	"WEEKEND",
}

func (d DayType) String() string {
	return dayTypeString[d]
}

func GetDayTypeForDateTime(dateTime time.Time) DayType {
	dayOfTheWeek := dateTime.Weekday()
	if dayOfTheWeek == time.Sunday || dayOfTheWeek == time.Saturday {
		return Weekend
	}
	return Weekday
}

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
