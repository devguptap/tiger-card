package icap

import (
	"tiger-card/cap/daily"
	"tiger-card/cap/pass"
	"tiger-card/cap/weekly"
	"tiger-card/trip"
)

// Cap interface define the method for all types of cap
type Cap interface {
	// GetCappedFare accept the trip details and initially calculated fare
	// and the validate it against the total fare and fare cap set
	IsCapLimitReached(t *trip.Trip, actualFare int) bool

	// Reset resets the cap
	Reset(t *trip.Trip)

	// UpdateTotalFare update teh total fare for a cap
	GetCappedFare(actualFare int) int

	// UpdateTotalFare update teh total fare for a cap
	UpdateCap(trip *trip.Trip)

	// UpdateTotalFare update teh total fare for a cap
	GetPass(trip *trip.Trip) *pass.Pass

	UpdateTotalFare(actualFare int)
}

// GetCaps returns all type of caps
func GetCaps() []Cap {
	return []Cap{
		new(daily.Daily),
		new(weekly.Weekly),
	}
}
