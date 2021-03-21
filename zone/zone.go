package zone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"tiger-card/logger"
)

type Id string
type Radius int

var zones map[Id]Radius

func InitZones() error {
	var err error
	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile("resources\\zones.json"); err == nil {
		zones = make(map[Id]Radius)
		if err = json.Unmarshal(fileBytes, zones); err == nil {
			logger.GetLogger().Info("[Config]. zones initialized successfully")
		}
	}

	return err
}

func GetZoneRadius(id string) (int, error) {
	if radius, ok := zones[Id(id)]; ok {
		return int(radius), nil
	} else {
		return -1, errors.New(fmt.Sprintf("invalid zone id : %s", id))
	}
}

func IsValidZone(id string) bool {
	_, ok := zones[Id(id)]
	return ok
}
