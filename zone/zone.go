package zone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"tiger-card/logger"
)

type Id string
type Radius int

var zoneRadiusMap map[Id]Radius

func InitZones() error {
	var err error
	var fileBytes []byte
	var resourceDirPath string
	if resourceDirPath = os.Getenv("TigerCardResourceDirPath"); resourceDirPath == "" {
		logger.GetLogger().Error("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
		return errors.New("unable to fetch resource directory path from env variable : TigerCardResourceDirPath")
	}
	if fileBytes, err = ioutil.ReadFile(resourceDirPath + string(os.PathSeparator) + "zones.json"); err == nil {
		zoneRadiusMap = make(map[Id]Radius)
		if err = json.Unmarshal(fileBytes, &zoneRadiusMap); err == nil {
			logger.GetLogger().Info("[Config]. zones initialized successfully")
		}
	}

	return err
}

func GetZoneRadius(id string) (int, error) {
	if radius, ok := zoneRadiusMap[Id(id)]; ok {
		return int(radius), nil
	} else {
		return -1, errors.New(fmt.Sprintf("invalid zone id : %s", id))
	}
}

func IsValidZone(id string) bool {
	_, ok := zoneRadiusMap[Id(id)]
	return ok
}

func GetZoneDistance(fromZone, toZone string) int {
	if zoneRadiusMap[Id(fromZone)] > zoneRadiusMap[Id(toZone)] {
		return int(zoneRadiusMap[Id(fromZone)] - zoneRadiusMap[Id(toZone)])
	} else {
		return int(zoneRadiusMap[Id(toZone)] - zoneRadiusMap[Id(fromZone)])
	}
}
