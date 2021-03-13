package peakhours

import (
	"tiger-card/constants"
	"tiger-card/reqresp"
	"time"
)

const (
	weekday = "weekday"
	weekend = "weekend"
)

type timeSlot struct {
	Start time.Time
	End   time.Time
}

var peakHoursMap = make(map[string][]*timeSlot)
var zone1ReturnJourneyOffPeakHour = make(map[string][]*timeSlot)

func InitPeakHours(weekdayTimeRange, weekendTimeRange map[string]string) error {
	var timeRange []*timeSlot
	var err error
	if timeRange, err = parseTimeRanges(weekdayTimeRange); err == nil {
		peakHoursMap[weekday] = timeRange
	} else {
		return err
	}

	if timeRange, err = parseTimeRanges(weekendTimeRange); err == nil {
		peakHoursMap[weekend] = timeRange
	} else {
		return err
	}
	return nil
}

func InitZone1ReturnJourneyOffPeeHours(weekdayTimeRange, weekendTimeRange map[string]string) error {
	var timeRange []*timeSlot
	var err error
	if timeRange, err = parseTimeRanges(weekdayTimeRange); err == nil {
		zone1ReturnJourneyOffPeakHour[weekday] = timeRange
	} else {
		return err
	}

	if timeRange, err = parseTimeRanges(weekendTimeRange); err == nil {
		zone1ReturnJourneyOffPeakHour[weekend] = timeRange
	} else {
		return err
	}
	return nil
}

func parseTimeRanges(fromToTimeMap map[string]string) ([]*timeSlot, error) {
	var timeRanges = make([]*timeSlot, 0, 0)
	var err error
	var startTime time.Time
	var endTime time.Time
	for startTimeString, endTimeString := range fromToTimeMap {
		if startTime, err = time.Parse(constants.TimeFormat, startTimeString); err == nil {
			if endTime, err = time.Parse(constants.TimeFormat, endTimeString); err == nil {
				timeRanges = append(timeRanges, &timeSlot{
					Start: startTime,
					End:   endTime,
				})
			}
		}

		if err != nil {
			break
		}
	}

	return timeRanges, err
}

func IsPeakHours(input *reqresp.TripData) bool {
	var isPeakHour bool
	isWeekend := isWeekend(input.DateTime)
	var timeSlotsToCompare []*timeSlot
	if isWeekend {
		timeSlotsToCompare = peakHoursMap[weekend]
	} else {
		timeSlotsToCompare = peakHoursMap[weekday]
	}

	for _, timeSlot := range timeSlotsToCompare {
		if isPeakHour = isWithinTimeRange(input.DateTime, timeSlot); isPeakHour == true {
			break
		}
	}

	if isPeakHour && input.ToZone > input.FromZone && input.FromZone == 1 && isReturnJourneyOffPeakHour(input.DateTime, isWeekend) {
		isPeakHour = false
	}

	return isPeakHour
}

func isReturnJourneyOffPeakHour(today time.Time, isWeekend bool) bool {
	var timeSlotsToCompare []*timeSlot

	if isWeekend {
		timeSlotsToCompare = zone1ReturnJourneyOffPeakHour[weekend]
	} else {
		timeSlotsToCompare = zone1ReturnJourneyOffPeakHour[weekday]
	}

	var isOffPeakHours bool

	for _, timeSlot := range timeSlotsToCompare {
		if isOffPeakHours = isWithinTimeRange(today, timeSlot); isOffPeakHours == true {
			break
		}
	}
	return isOffPeakHours
}

func isWeekend(today time.Time) bool {
	dayOfTheWeek := today.Weekday()
	if dayOfTheWeek == time.Sunday || dayOfTheWeek == time.Saturday {
		return true
	}
	return false
}

func isWithinTimeRange(today time.Time, timeRange *timeSlot) bool {
	todayHr := today.Hour()
	todayMin := today.Minute()
	startHr := timeRange.Start.Hour()
	startMin := timeRange.Start.Minute()
	endHr := timeRange.End.Hour()
	endMin := timeRange.End.Minute()
	if (todayHr < startHr || todayHr > endHr) ||
		(todayHr == startHr && todayMin < startMin) ||
		(todayHr == endHr && todayMin > endMin) {
		return false
	}
	return true
}
