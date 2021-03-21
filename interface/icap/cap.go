package icap

import (
	"tiger-card/pass/daily"
	"tiger-card/pass/weekly"
	"tiger-card/trip"
	"time"
)

// Cap interface define the method for all types of cap
type Cap interface {
	// GetCappedFare accept the trip details and initially calculated fare
	// and the validate it against the total fare and fare cap set
	GetCappedFare(trip *trip.Trip, actualFare int) int

	// Reset resets the cap
	Reset(dateTime time.Time)

	// UpdateTotalFare update teh total fare for a cap
	UpdateTotalFare(fare int)
}

// GetCaps returns all type of caps
func GetCaps() []Cap {
	return []Cap{
		new(daily.Daily),
		new(weekly.Weekly),
	}
}
