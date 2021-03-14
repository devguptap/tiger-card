package icap

import (
	"tiger-card/cap/daily"
	"tiger-card/cap/weekly"
	"tiger-card/trip"
	"time"
)

type Cap interface {
	GetCappedFare(trip *trip.Trip, actualFare int) int
	Reset(dateTime time.Time)
	UpdateTotalFare(fare int)
}

func GetCaps() []Cap {
	return []Cap{
		new(daily.Daily),
		new(weekly.Weekly),
	}
}
