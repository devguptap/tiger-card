package farecalculator

import (
	"tiger-card/caps"
	"tiger-card/config"
	"tiger-card/fare"
	"tiger-card/peakhours"
	"tiger-card/reqresp"
	"time"
)

type capAndTotalFareObj struct {
	fareCap     int
	zoneDiff    int
	totalFare   int
	lastUpdated time.Time
}

var dailyData = new(capAndTotalFareObj)
var weeklyData = new(capAndTotalFareObj)

var startOfTheWeekDayDiff = map[time.Weekday]int{
	time.Monday:    0,
	time.Tuesday:   1,
	time.Wednesday: 2,
	time.Thursday:  3,
	time.Friday:    4,
	time.Saturday:  5,
	time.Sunday:    6,
}

func CalculateFare(trips []*reqresp.TripData) int {
	reset(trips[0])
	var totalFare int
	for _, trip := range trips {
		checkAndUpdateWeeklyCap(trip)
		checkAndUpdateDailyCap(trip)
		tripFare := fare.GetFare(trip.FromZone, trip.ToZone, peakhours.IsPeakHours(trip))
		effectiveWeeklyCap := getEffectiveWeeklyCap()
		effectiveDailyCap := getEffectiveDailyCap()
		if effectiveWeeklyCap < tripFare {
			tripFare = effectiveWeeklyCap
		} else if effectiveDailyCap < tripFare {
			tripFare = effectiveDailyCap
		}
		totalFare += tripFare
		dailyData.totalFare += tripFare
		weeklyData.totalFare += tripFare
	}
	return totalFare
}

func reset(trip *reqresp.TripData) {
	dailyData.fareCap, dailyData.totalFare, weeklyData.fareCap, weeklyData.totalFare = 0, 0, 0, 0
	dailyData.zoneDiff, weeklyData.zoneDiff = -1, -1
	dailyData.lastUpdated, weeklyData.lastUpdated = trip.DateTime, trip.DateTime
}

func resetDailyLimit(trip *reqresp.TripData) {
	dailyData.fareCap, dailyData.totalFare = 0, 0
	dailyData.lastUpdated = trip.DateTime
	dailyData.zoneDiff = -1
}

func getZonDiff(z1, z2 int) int {
	if z1 > z2 {
		return z1 - z2
	} else {
		return z2 - z1
	}
}

func checkAndUpdateDailyCap(trip *reqresp.TripData) {
	if dailyData.lastUpdated.Weekday() != trip.DateTime.Weekday() {
		resetDailyLimit(trip)
	}
	currZoneDiff := getZonDiff(trip.FromZone, trip.ToZone)

	if currZoneDiff > dailyData.zoneDiff {
		dailyData.zoneDiff = currZoneDiff
		dailyData.fareCap = caps.GetDailyCap(trip.FromZone, trip.ToZone)
	}
}

func checkAndUpdateWeeklyCap(trip *reqresp.TripData) {
	if weeklyData.lastUpdated.Before(getStartOfTheWeekDateTime(trip.DateTime)) {
		reset(trip)
	}
	currZoneDiff := getZonDiff(trip.FromZone, trip.ToZone)

	if currZoneDiff > weeklyData.zoneDiff {
		weeklyData.zoneDiff = currZoneDiff
		weeklyData.fareCap = caps.GetWeeklyCap(trip.FromZone, trip.ToZone)
	}
}

func getStartOfTheWeekDateTime(today time.Time) time.Time {
	currDay := today.Weekday()
	startOfTheWeek := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, config.ISTLocation).AddDate(0, 0, -(startOfTheWeekDayDiff[currDay]))
	return startOfTheWeek
}

func getEffectiveWeeklyCap() int {
	return weeklyData.fareCap - weeklyData.totalFare
}

func getEffectiveDailyCap() int {
	return dailyData.fareCap - dailyData.totalFare
}
