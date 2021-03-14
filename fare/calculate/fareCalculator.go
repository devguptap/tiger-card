package calculate

import (
	"tiger-card/fare"
	"tiger-card/interface/icap"
	"tiger-card/trip"
	"time"
)

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

func getCappedFare(actualFare int, trip *trip.Trip, caps []icap.Cap) int {
	for _, c := range caps {
		if cappedFare := c.GetCappedFare(trip, actualFare); cappedFare < actualFare {
			actualFare = cappedFare
		}
	}
	return actualFare
}

func updateCapsTotalFare(caps []icap.Cap, actualfare int) {
	for _, c := range caps {
		c.UpdateTotalFare(actualfare)
	}
}

func resetAllCaps(caps []icap.Cap, dateTime time.Time) {
	for _, c := range caps {
		c.Reset(dateTime)
	}
}
