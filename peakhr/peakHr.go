package peakhr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"tiger-card/config"
	"tiger-card/logger"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

var peakHrMap map[time.Weekday][]*timeRange
var returnTripOffPeakHrMap map[zone.Id]map[zone.Id]map[time.Weekday][]*timeRange

type timeRange struct {
	fromTime time.Time
	toTime   time.Time
}

type configTimeObj struct {
	FromTime string `json:"fromTime"`
	ToTime   string `json:"toTime"`
}

// IsPeakHours returns whether this trip falls in peak time
func IsPeakHours(trip *trip.Trip) bool {
	peakHrRange := peakHrMap[trip.DateTime.Weekday()]
	var result bool
	for _, timeRange := range peakHrRange {
		if result = isDateTimeWithinTimeRange(trip.DateTime, timeRange); result == true {
			break
		}
	}

	if result && isReturnTripOffPeakHr(trip) {
		result = false
	}
	return result
}

func isReturnTripOffPeakHr(trip *trip.Trip) bool {
	var result bool
	if _, ok := returnTripOffPeakHrMap[zone.Id(trip.FromZone)]; ok {
		if _, ok := returnTripOffPeakHrMap[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)]; ok {
			timeRanges := returnTripOffPeakHrMap[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)][trip.DateTime.Weekday()]
			for _, timeRange := range timeRanges {
				if result = isDateTimeWithinTimeRange(trip.DateTime, timeRange); result == true {
					break
				}
			}
		}
	}
	return result
}

// isDateTimeWithinTimeRange check whether the given dateTime falls in timeRange
func isDateTimeWithinTimeRange(dateTime time.Time, timeRange *timeRange) bool {
	todayHr := dateTime.Hour()
	todayMin := dateTime.Minute()
	startHr := timeRange.fromTime.Hour()
	startMin := timeRange.fromTime.Minute()
	endHr := timeRange.toTime.Hour()
	endMin := timeRange.toTime.Minute()
	if (todayHr < startHr || todayHr > endHr) ||
		(todayHr == startHr && todayMin < startMin) ||
		(todayHr == endHr && todayMin > endMin) {
		return false
	}
	return true
}

func InitPeakHrMap() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "peakHourConfig.json"); err == nil {
		peakHrConfig := make(map[string][]*configTimeObj)
		if err = json.Unmarshal(fileBytes, &peakHrConfig); err == nil {
			if peakHrMap, err = parseConfigForDayTimeRangeMap(peakHrConfig); err == nil {
				logger.GetLogger().Info("Peak Hour Config initialized successfully")
			}
		}
	}
	return err
}

func InitReturnTripOffPeakHrMap() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "returnTripOffPeekHourConfig.json"); err == nil {
		offPeakHrConfig := make(map[zone.Id]map[zone.Id]map[string][]*configTimeObj)
		if err = json.Unmarshal(fileBytes, &offPeakHrConfig); err == nil {
			returnTripOffPeakHrMap = make(map[zone.Id]map[zone.Id]map[time.Weekday][]*timeRange)
			for fromZone, toZoneMap := range offPeakHrConfig {
				for toZone, offPeakHrConfigMap := range toZoneMap {
					var offPeakHrMap = make(map[time.Weekday][]*timeRange)
					if offPeakHrMap, err = parseConfigForDayTimeRangeMap(offPeakHrConfigMap); err == nil {
						returnTripOffPeakHrMap[fromZone] = make(map[zone.Id]map[time.Weekday][]*timeRange)
						returnTripOffPeakHrMap[fromZone][toZone] = offPeakHrMap
					} else {
						return err
					}
				}
			}

		}
	}
	return err
}

func parseConfigForDayTimeRangeMap(parsedConfig map[string][]*configTimeObj) (map[time.Weekday][]*timeRange, error) {
	var err error
	dayTimeRangeMap := make(map[time.Weekday][]*timeRange)
	for dayString, timeRanges := range parsedConfig {
		var day time.Weekday
		switch dayString {
		case time.Monday.String():
			day = time.Monday
		case time.Tuesday.String():
			day = time.Tuesday
		case time.Wednesday.String():
			day = time.Wednesday
		case time.Thursday.String():
			day = time.Thursday
		case time.Friday.String():
			day = time.Friday
		case time.Saturday.String():
			day = time.Saturday
		case time.Sunday.String():
			day = time.Sunday
		default:
			err = errors.New(fmt.Sprintf("Invalid day string : %s in peak hr config", dayString))
		}

		if err == nil {
			var dayWiseTimeRange []*timeRange
			if dayWiseTimeRange, err = getTimeRanges(timeRanges); err == nil {
				dayTimeRangeMap[day] = dayWiseTimeRange
			} else {
				break
			}
		} else {
			break
		}

	}

	return dayTimeRangeMap, err
}

func getTimeRanges(configTimeRanges []*configTimeObj) ([]*timeRange, error) {
	dayWiseTimeRange := make([]*timeRange, 0, 0)
	var err error
	for _, tr := range configTimeRanges {
		var fromTime time.Time
		if fromTime, err = time.ParseInLocation("15:04", tr.FromTime, config.ISTLocation); err == nil {
			var toTime time.Time
			if toTime, err = time.ParseInLocation("15:04", tr.ToTime, config.ISTLocation); err == nil {
				dayWiseTimeRange = append(dayWiseTimeRange, &timeRange{
					fromTime: fromTime,
					toTime:   toTime,
				})
			}
		}
		if err != nil {
			break
		}
	}
	return dayWiseTimeRange, err
}
