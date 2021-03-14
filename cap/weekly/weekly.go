package weekly

import (
	"tiger-card/cap/util"
	"tiger-card/config"
	"tiger-card/trip"
	"time"
)

var dayDiffFromStartOfTheWeek = map[time.Weekday]int{
	time.Monday:    0,
	time.Tuesday:   1,
	time.Wednesday: 2,
	time.Thursday:  3,
	time.Friday:    4,
	time.Saturday:  5,
	time.Sunday:    6,
}

var weeklyCap [][]int

type Weekly struct {
	capAmount    int
	zoneDistance int
	totalFare    int
	lastUpdated  time.Time
}

func (w *Weekly) Reset(dateTime time.Time) {
	w.capAmount, w.totalFare = 0, 0
	w.zoneDistance = -1
	w.lastUpdated = dateTime
}

func (w *Weekly) GetCappedFare(trip *trip.Trip, actualFare int) int {
	w.updateCap(trip)
	if w.capAmount-w.totalFare < actualFare {
		return w.capAmount - w.totalFare
	} else {
		return actualFare
	}
}

func (w *Weekly) updateCap(trip *trip.Trip) {
	if w.lastUpdated.Before(getStartOfTheWeekDateTime(trip.DateTime)) {
		w.Reset(trip.DateTime)
	}

	currZoneDistance := util.GetZoneDistance(trip.FromZone, trip.ToZone)
	if currZoneDistance > w.zoneDistance {
		w.zoneDistance = currZoneDistance
		w.capAmount = weeklyCap[trip.FromZone.GetId()-1][trip.ToZone.GetId()-1]
	}
}

func (w *Weekly) UpdateTotalFare(fare int) {
	w.totalFare += fare
}

func Init(capMatrix [][]int) {
	weeklyCap = capMatrix
}

func getStartOfTheWeekDateTime(today time.Time) time.Time {
	currDay := today.Weekday()
	startOfTheWeek := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, config.ISTLocation).AddDate(0, 0, -(dayDiffFromStartOfTheWeek[currDay]))
	return startOfTheWeek
}
