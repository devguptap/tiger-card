package daily

import (
	"encoding/json"
	"io/ioutil"
	"tiger-card/logger"
	"tiger-card/pass/util"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

// dailyCap contains the daily cap matrix for zone combination for a particular test suite. See example below
// [100	 120]
// [120	  80]
type dailyCapAmount int

var dailyCapMap map[zone.Id]map[zone.Id]dailyCapAmount

type Daily struct {
	capAmount  int
	dailyTotal int
	expiry     time.Time
	isActive   bool
}

func (d *Daily) Reset(dateTime time.Time) {
	d.capAmount = 0
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

func InitDailyCap() error {
	var err error
	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile("resources\\zones.json"); err == nil {
		zone = make(map[ZoneId]ZoneRadius)
		if err = json.Unmarshal(fileBytes, zone); err == nil {
			logger.GetLogger().Info("[Config]. zones initialized successfully")
		}
	}

	return err
}
