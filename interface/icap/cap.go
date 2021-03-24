package icap

import (
	"tiger-card/cap/daily"
	"tiger-card/cap/pass"
	"tiger-card/cap/weekly"
	"tiger-card/trip"
)

// Cap interface define the method for all type of cap (daily / weekly)
type Cap interface {
	// IsCapLimitReached accept the actual fare and check if this cap limit is reached.
	IsCapLimitReached(actualFare int) bool

	// Reset resets the cap
	Reset(t *trip.Trip)

	// GetCappedFare return the applicable fare for this cap.
	GetCappedFare(actualFare int) int

	// UpdateCap update the cap variable like totalFare etc for this cap
	UpdateCap(trip *trip.Trip)

	// GetPass return the pass with an expiry for this cap
	GetPass(trip *trip.Trip) *pass.Pass

	// UpdateTotalFare update the total fare for this cap
	UpdateTotalFare(actualFare int)
}

// GetCaps returns all type of caps
func GetCaps() []Cap {
	return []Cap{
		new(daily.Daily),
		new(weekly.Weekly),
	}
}
