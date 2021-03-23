package weekly

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"tiger-card/cap/pass"
	"tiger-card/logger"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

// weeklyCap contains the weekly cap matrix for zone combination for a particular test suite. See example below
// [500	 600]
// [600	 400]

var weeklyCapMap map[zone.Id]map[zone.Id]int

type Weekly struct {
	capAmount    int
	weeklyTotal  int
	zoneDistance int
	expiry       time.Time
}

func (w *Weekly) Reset(t *trip.Trip) {
	w.capAmount = weeklyCapMap[zone.Id(t.FromZone)][zone.Id(t.ToZone)]
	w.zoneDistance = zone.GetZoneDistance(t.FromZone, t.ToZone)
	w.expiry = getWeeklyCapExpiry(t.DateTime)
	w.weeklyTotal = 0
}

func (w *Weekly) IsCapLimitReached(t *trip.Trip, actualFare int) bool {
	if zone.GetZoneDistance(t.FromZone, t.ToZone) > w.zoneDistance {
		w.zoneDistance = zone.GetZoneDistance(t.FromZone, t.ToZone)
		w.capAmount = weeklyCapMap[zone.Id(t.FromZone)][zone.Id(t.ToZone)]
	}

	if w.capAmount-w.weeklyTotal < actualFare {
		return true
	} else {
		return false
	}
}

func (w *Weekly) GetCappedFare(actualFare int) int {
	if w.capAmount-w.weeklyTotal < actualFare {
		return w.capAmount - w.weeklyTotal
	} else {
		return actualFare
	}
}

func (w *Weekly) UpdateCap(trip *trip.Trip) {
	if w.expiry.After(trip.DateTime) {
		currZoneDistance := zone.GetZoneDistance(trip.ToZone, trip.ToZone)
		if currZoneDistance > w.zoneDistance {
			w.zoneDistance = currZoneDistance
			w.capAmount = weeklyCapMap[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)]
		}
	} else {
		w.Reset(trip)
	}
}

func (w *Weekly) UpdateTotalFare(actualFare int) {
	w.weeklyTotal += actualFare
}

func InitWeeklyCap() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "weeklyCap.json"); err == nil {
		weeklyCapMap = make(map[zone.Id]map[zone.Id]int)
		if err = json.Unmarshal(fileBytes, &weeklyCapMap); err == nil {
			logger.GetLogger().Info("[Config]. weekly cap initialized successfully")
		}
	}

	return err
}

func getWeeklyCapExpiry(t time.Time) time.Time {
	dayOfTheWeek := int(t.Weekday())
	dayToAdd := (7-dayOfTheWeek)%7 + 1
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, dayToAdd)
}

func (w *Weekly) GetPass(trip *trip.Trip) *pass.Pass {
	return pass.NewPass("weekly", getWeeklyCapExpiry(trip.DateTime))
}
