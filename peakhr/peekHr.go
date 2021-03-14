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

type timeRange struct {
	From *timeStamp
	To   *timeStamp
}

type timeStamp struct {
	Hour int
	Min  int
}

type peakHrRule struct {
	PeakHrMap map[zone.Zone]map[zone.Zone]map[day.DayType][]*timeRange
}

var peakHr = &peakHrRule{PeakHrMap: make(map[zone.Zone]map[zone.Zone]map[day.DayType][]*timeRange)}

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

func getToZonePeakHrMap(toZonePeakHrMap map[string]map[string][]string) map[zone.Zone]map[day.DayType][]*timeRange {
	var result = make(map[zone.Zone]map[day.DayType][]*timeRange)
	for toZoneId, dayTypePeakHrMap := range toZonePeakHrMap {
		toZone := zone.NewZoneForZoneId(toZoneId)
		result[toZone] = getDayTypeTimeRangeMap(dayTypePeakHrMap)
	}
	return result
}

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
