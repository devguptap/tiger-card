package peakhr

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tiger-card/enum/day"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

// timeRange contains the peak hour time range
type timeRange struct {
	From *timeStamp
	To   *timeStamp
}

// timeStamp contains the peak hour hour and min value
type timeStamp struct {
	Hour int
	Min  int
}

// peakHrRule contains all the peak hour set for a combination of zones
type peakHrRule struct {
	PeakHrMap map[zone.Zone]map[zone.Zone]map[day.DayType][]*timeRange
}

// peakHr is the variable which contains all the peak hour set for a combination of zones for a test suite
var peakHr = &peakHrRule{PeakHrMap: make(map[zone.Zone]map[zone.Zone]map[day.DayType][]*timeRange)}

// IsPeakHours returns whether this trip falls in peak time
func IsPeakHours(trip *trip.Trip) bool {
	var isPeakHour bool
	for _, timeRange := range getTimeRangeForTrip(trip) {
		if isDateTimeWithinTimeRange(trip.DateTime, timeRange) {
			isPeakHour = true
			break
		}
	}

	return isPeakHour
}

// isDateTimeWithinTimeRange check whether the given dateTime falls in timeRange
func isDateTimeWithinTimeRange(dateTime time.Time, timeRange *timeRange) bool {
	todayHr := dateTime.Hour()
	todayMin := dateTime.Minute()
	startHr := timeRange.From.Hour
	startMin := timeRange.From.Min
	endHr := timeRange.To.Hour
	endMin := timeRange.To.Min
	if (todayHr < startHr || todayHr > endHr) ||
		(todayHr == startHr && todayMin < startMin) ||
		(todayHr == endHr && todayMin > endMin) {
		return false
	}
	return true
}

// getTimeRangeForTrip accept the trip and return the list of peak time range configured for the zone combination
func getTimeRangeForTrip(trip *trip.Trip) []*timeRange {
	if toZoneMap, ok := peakHr.PeakHrMap[trip.FromZone]; ok {
		if dayRangeMap, ok := toZoneMap[trip.ToZone]; ok {
			if timeRangeList, ok := dayRangeMap[day.GetDayTypeForDateTime(trip.DateTime)]; ok {
				return timeRangeList
			}
		}
	}
	return nil
}

// InitializePeakHour Read the config file bytes and initialize the peakHr varuable
func InitializePeakHour(fileBytes []byte) {
	var configFileObj = make(map[string]map[string]map[string][]string)
	var err error
	if err = json.Unmarshal(fileBytes, &configFileObj); err == nil {
		for fromZoneId, toZonePeakHrMap := range configFileObj {
			fromZone := zone.NewZoneForZoneId(fromZoneId)
			peakHr.PeakHrMap[fromZone] = getToZonePeakHrMap(toZonePeakHrMap)
		}
	} else {
		panic(err)
	}

}

// getToZonePeakHrMap is a helper function for InitializePeakHour
func getToZonePeakHrMap(toZonePeakHrMap map[string]map[string][]string) map[zone.Zone]map[day.DayType][]*timeRange {
	var result = make(map[zone.Zone]map[day.DayType][]*timeRange)
	for toZoneId, dayTypePeakHrMap := range toZonePeakHrMap {
		toZone := zone.NewZoneForZoneId(toZoneId)
		result[toZone] = getDayTypeTimeRangeMap(dayTypePeakHrMap)
	}
	return result
}

// getDayTypeTimeRangeMap is a helper function for getToZonePeakHrMap
func getDayTypeTimeRangeMap(dayTypePeakHrMap map[string][]string) map[day.DayType][]*timeRange {
	var result = make(map[day.DayType][]*timeRange)
	for dt, peakHrRange := range dayTypePeakHrMap {
		dayType := day.GetDayTypeForString(dt)
		var timeRangeSlice = make([]*timeRange, 0, 0)
		for _, peakHr := range peakHrRange {
			timeRangeSlice = append(timeRangeSlice, parseTimeRangeString(peakHr))
		}
		result[dayType] = timeRangeSlice
	}
	return result
}

// parseTimeRangeString parses the time range string. Time range should be in format HH:mm-HH:mm
func parseTimeRangeString(tr string) *timeRange {
	var timeRangeObj *timeRange
	trSlice := strings.Split(tr, "-")
	if len(trSlice) == 2 {
		timeRangeObj = &timeRange{
			From: parseTimeStamp(trSlice[0]),
			To:   parseTimeStamp(trSlice[1]),
		}
	} else {
		panic(errors.New(fmt.Sprintf("invalid time range  : %s", tr)))
	}
	return timeRangeObj
}

// parseTimeStamp parses the timestamp. Timestamp should be in format : HH:mm
func parseTimeStamp(hm string) *timeStamp {
	var timeStampObj *timeStamp
	var err error
	hmSlice := strings.Split(hm, ":")
	if len(hmSlice) == 2 {
		var hr int
		if hr, err = strconv.Atoi(strings.TrimSpace(hmSlice[0])); err == nil {
			var min int
			if min, err = strconv.Atoi(strings.TrimSpace(hmSlice[1])); err == nil {
				if hr >= 0 && hr <= 23 && min >= 0 && min <= 59 {
					timeStampObj = &timeStamp{
						Hour: hr,
						Min:  min,
					}
				} else {
					err = errors.New(fmt.Sprintf("Invalid timestamp : %s", hm))
				}
			}
		}
	} else {
		err = errors.New(fmt.Sprintf("invalid timestamp : %s", hm))
	}

	if err != nil {
		panic(err)
	}
	return timeStampObj
}
