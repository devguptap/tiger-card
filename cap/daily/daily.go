package daily

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

// dailyCap contains the daily cap matrix for zone combination for a particular test suite. See example below
// [100	 120]
// [120	  80]

var dailyCapMap map[zone.Id]map[zone.Id]int

type Daily struct {
	capAmount    int
	dailyTotal   int
	zoneDistance int
	expiry       time.Time
}

func (d *Daily) Reset(t *trip.Trip) {
	d.capAmount = dailyCapMap[zone.Id(t.FromZone)][zone.Id(t.ToZone)]
	d.zoneDistance = zone.GetZoneDistance(t.FromZone, t.ToZone)
	d.expiry = getDailyCapExpiry(t.DateTime)
	d.dailyTotal = 0
}

func (d *Daily) IsCapLimitReached(t *trip.Trip, actualFare int) bool {
	if d.capAmount-d.dailyTotal < actualFare {
		return true
	} else {
		return false
	}
}

func (d *Daily) GetCappedFare(actualFare int) int {
	if d.capAmount-d.dailyTotal < actualFare {
		return d.capAmount - d.dailyTotal
	} else {
		return actualFare
	}
}

func (d *Daily) UpdateCap(trip *trip.Trip) {
	if d.expiry.After(trip.DateTime) {
		currZoneDistance := zone.GetZoneDistance(trip.ToZone, trip.ToZone)
		if currZoneDistance > d.zoneDistance {
			d.zoneDistance = currZoneDistance
			d.capAmount = dailyCapMap[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)]
		}
	} else {
		d.Reset(trip)
	}
}

func (d *Daily) UpdateTotalFare(actualFare int) {
	d.dailyTotal += actualFare
}

func InitDailyCap() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "dailyCap.json"); err == nil {
		dailyCapMap = make(map[zone.Id]map[zone.Id]int)
		if err = json.Unmarshal(fileBytes, &dailyCapMap); err == nil {
			logger.GetLogger().Info("[Config]. daily cap initialized successfully")
		}
	}

	return err
}

func getDailyCapExpiry(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day()+1, 0, 0, 0, 0, timestamp.Location())
}

func (d *Daily) GetPass(trip *trip.Trip) *pass.Pass {
	return pass.NewPass("daily", getDailyCapExpiry(trip.DateTime))
}
