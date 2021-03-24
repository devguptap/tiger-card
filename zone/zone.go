package zone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"tiger-card/logger"
)

// Id represent a zone id like Z1, Z2 etc
type Id string

// Radius represent the radius of the zone from the center of the city
type Radius int

// zoneRadiusMap contains the map of zone and its radius from center
var zoneRadiusMap map[Id]Radius

// InitZones parses the zones.json config file and initializes teh zone radius map.
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

// GetZoneRadius return the zone radius for zone id
func GetZoneRadius(id string) (int, error) {
	if radius, ok := zoneRadiusMap[Id(id)]; ok {
		return int(radius), nil
	} else {
		return -1, errors.New(fmt.Sprintf("invalid zone id : %s", id))
	}
}

// IsValidZone check and validate if the zone is is a valid zone id.
func IsValidZone(id string) bool {
	_, ok := zoneRadiusMap[Id(id)]
	return ok
}

// GetZoneDistance return the distance between fromZone and toZone based on their radius from the center of the city
func GetZoneDistance(fromZone, toZone string) int {
	if zoneRadiusMap[Id(fromZone)] > zoneRadiusMap[Id(toZone)] {
		return int(zoneRadiusMap[Id(fromZone)] - zoneRadiusMap[Id(toZone)])
	} else {
		return int(zoneRadiusMap[Id(toZone)] - zoneRadiusMap[Id(fromZone)])
	}
}
