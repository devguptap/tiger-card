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

// weeklyCapMap represent the mapping of weekly cap between zone
var weeklyCapMap map[zone.Id]map[zone.Id]int

// Weekly struct contains the weekly cap limit and weekly total for a tiger-card
type Weekly struct {
	capAmount    int
	weeklyTotal  int
	zoneDistance int
	expiry       time.Time
}

// Reset reset the weekly cap for the first journey of the week
func (w *Weekly) Reset(t *trip.Trip) {
	w.capAmount = weeklyCapMap[zone.Id(t.FromZone)][zone.Id(t.ToZone)]
	w.zoneDistance = zone.GetZoneDistance(t.FromZone, t.ToZone)
	w.expiry = getWeeklyCapExpiry(t.DateTime)
	w.weeklyTotal = 0
}

// IsCapLimitReached check if weekly total fare + current trip fare is exceeding the cap limit
func (w *Weekly) IsCapLimitReached(actualFare int) bool {
	if w.capAmount-w.weeklyTotal < actualFare {
		return true
	} else {
		return false
	}
}

// GetCappedFare return the modified fare if cap limit reached else return the actualFare
func (w *Weekly) GetCappedFare(actualFare int) int {
	if w.capAmount-w.weeklyTotal < actualFare {
		return w.capAmount - w.weeklyTotal
	} else {
		return actualFare
	}
}

// UpdateCap update the weekly cap object for given trip. It reset the cap if trip is for new week
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

// UpdateTotalFare update the weekly total fare with current fare
func (w *Weekly) UpdateTotalFare(actualFare int) {
	w.weeklyTotal += actualFare
}

// InitWeeklyCap initialize the weeklyCapMap with the provided config
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

// getWeeklyCapExpiry set the expiry of weekly cap object for the end of the week
func getWeeklyCapExpiry(t time.Time) time.Time {
	dayOfTheWeek := int(t.Weekday())
	dayToAdd := (7-dayOfTheWeek)%7 + 1
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, dayToAdd)
}

// GetPass returns a weekly pass. If pass achieved, then ride will be free till the expiry of the pass.
func (w *Weekly) GetPass(trip *trip.Trip) *pass.Pass {
	return pass.NewPass("weekly", getWeeklyCapExpiry(trip.DateTime))
}
