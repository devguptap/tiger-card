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

// dailyCapMap represent the mapping of daily cap between zone
var dailyCapMap map[zone.Id]map[zone.Id]int

// Daily struct contains the daily cap limit and daily total for a tiger-card
type Daily struct {
	capAmount    int
	dailyTotal   int
	zoneDistance int
	expiry       time.Time
}

// Reset reset the daily cap for the first journey of the day
func (d *Daily) Reset(t *trip.Trip) {
	d.capAmount = dailyCapMap[zone.Id(t.FromZone)][zone.Id(t.ToZone)]
	d.zoneDistance = zone.GetZoneDistance(t.FromZone, t.ToZone)
	d.expiry = getDailyCapExpiry(t.DateTime)
	d.dailyTotal = 0
}

// IsCapLimitReached check if daily total fare + current trip fare is exceeding the cap limit
func (d *Daily) IsCapLimitReached(actualFare int) bool {
	if d.capAmount-d.dailyTotal < actualFare {
		return true
	} else {
		return false
	}
}

// GetCappedFare return the modified fare if cap limit reached else return the actualFare
func (d *Daily) GetCappedFare(actualFare int) int {
	if d.capAmount-d.dailyTotal < actualFare {
		return d.capAmount - d.dailyTotal
	} else {
		return actualFare
	}
}

// UpdateCap update the daily cap object for given trip. It reset the cap if trip is for new day
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

// UpdateTotalFare update the daily total fare with current fare
func (d *Daily) UpdateTotalFare(actualFare int) {
	d.dailyTotal += actualFare
}

// InitDailyCap initialize the dailyCapMap with the provided config
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

// getDailyCapExpiry set the expiry of daily cap object for the end of the day
func getDailyCapExpiry(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day()+1, 0, 0, 0, 0, timestamp.Location())
}

// GetPass returns a daily pass. If pass achieved, then ride will be free till the expiry of the pass.
func (d *Daily) GetPass(trip *trip.Trip) *pass.Pass {
	return pass.NewPass("daily", getDailyCapExpiry(trip.DateTime))
}
