package fare

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"tiger-card/logger"
	"tiger-card/peakhr"
	"tiger-card/trip"
	"tiger-card/zone"
)

type fareType string

const (
	peakFare    fareType = "peakFare"
	offPeakFare fareType = "offPeakFare"
)

var fares = make(map[zone.Id]map[zone.Id]map[fareType]int)

// InitFare initializes the zone wise fare
func InitFare() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "fare.json"); err == nil {
		if err = json.Unmarshal(fileBytes, &fares); err == nil {
			logger.GetLogger().Info("[Config]. fare initialized successfully")
		}
	}
	return err
}

// GetFare return the fare applicable for a single trip based on from and to zone
func GetFare(trip *trip.Trip) int {
	if peakhr.IsPeakHours(trip) {
		return fares[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)][peakFare]
	} else {
		return fares[zone.Id(trip.FromZone)][zone.Id(trip.ToZone)][offPeakFare]
	}
}
