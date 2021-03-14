package daily

import (
	"tiger-card/cap/util"
	"tiger-card/trip"
	"time"
)

// dailyCap contains the daily cap matrix for zone combination for a particular test suite. See example below
// [100	 120]
// [120	  80]
var dailyCap [][]int

type Daily struct {
	capAmount    int
	zoneDistance int
	totalFare    int
	lastUpdated  time.Time
}

func (d *Daily) Reset(dateTime time.Time) {
	d.capAmount, d.totalFare = 0, 0
	d.zoneDistance = -1
	d.lastUpdated = dateTime
}

func (d *Daily) GetCappedFare(trip *trip.Trip, actualFare int) int {
	d.updateCap(trip)
	if d.capAmount-d.totalFare < actualFare {
		return d.capAmount - d.totalFare
	} else {
		return actualFare
	}
}

func (d *Daily) updateCap(trip *trip.Trip) {
	if d.lastUpdated.Weekday() != trip.DateTime.Weekday() {
		d.Reset(trip.DateTime)
	}
	currZoneDistance := util.GetZoneDistance(trip.FromZone, trip.ToZone)

	if currZoneDistance > d.zoneDistance {
		d.zoneDistance = currZoneDistance
		d.capAmount = dailyCap[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	}
}

func (d *Daily) UpdateTotalFare(fare int) {
	d.totalFare += fare
}

func Init(capMatrix [][]int) {
	dailyCap = capMatrix
}
