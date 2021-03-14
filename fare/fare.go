package fare

import (
	"tiger-card/peakhr"
	"tiger-card/trip"
)

type Fare struct {
	peakHr, offPeakHr [][]int
}

var fareObj = new(Fare)

func Init(peakHrFareMatrix, offPeakHrFareMatrix [][]int) {
	fareObj.peakHr = peakHrFareMatrix
	fareObj.offPeakHr = offPeakHrFareMatrix
}

func GetFare(trip *trip.Trip) int {
	if peakhr.IsPeakHours(trip) {
		return fareObj.peakHr[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	} else {
		return fareObj.offPeakHr[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	}
}
