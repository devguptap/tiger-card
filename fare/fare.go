package fare

import (
	"tiger-card/peakhr"
	"tiger-card/trip"
)

type Fare struct {
	peakHr, offPeakHr [][]int
}

// fareObj contains all the peak and off-peak fare for a particular test suit
// in a form of matrix each row and column corresponding to the zones in the city
var fareObj = new(Fare)

// Init initializes the fare object
func Init(peakHrFareMatrix, offPeakHrFareMatrix [][]int) {
	fareObj.peakHr = peakHrFareMatrix
	fareObj.offPeakHr = offPeakHrFareMatrix
}

// GetFare return the fare applicable for a single trip
func GetFare(trip *trip.Trip) int {
	if peakhr.IsPeakHours(trip) {
		return fareObj.peakHr[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	} else {
		return fareObj.offPeakHr[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	}
}
