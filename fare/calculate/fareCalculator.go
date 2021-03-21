package calculate

import (
	"tiger-card/fare"
	"tiger-card/interface/icap"
	"tiger-card/trip"
	"time"
)

// FareCalculator accepts the list of trips and return the total fare for all the trip
func FareCalculator(trips []*trip.Trip) int {

	if len(trips) == 0 {
		return 0
	}
	caps := icap.GetCaps()
	resetAllCaps(caps, trips[0].DateTime)

	var totalFare int
	for _, t := range trips {
		actualFare := fare.GetFare(t)
		actualFare = getCappedFare(actualFare, t, caps)
		updateCapsTotalFare(caps, actualFare)
		totalFare += actualFare
	}
	return totalFare
}

// getCappedFare validate the actual fare against all the caps and return the discounted fare.
func getCappedFare(actualFare int, trip *trip.Trip, caps []icap.Cap) int {
	for _, c := range caps {
		if cappedFare := c.GetCappedFare(trip, actualFare); cappedFare < actualFare {
			actualFare = cappedFare
		}
	}
	return actualFare
}

// updateCapsTotalFare update the total fare for the list of journey for all the caps
func updateCapsTotalFare(caps []icap.Cap, actualFare int) {
	for _, c := range caps {
		c.UpdateTotalFare(actualFare)
	}
}

// resetAllCaps reset all the caps
func resetAllCaps(caps []icap.Cap, dateTime time.Time) {
	for _, c := range caps {
		c.Reset(dateTime)
	}
}
