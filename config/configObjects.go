package config

import (
	"errors"
	"fmt"
)

type appConfig struct {
	zoneRadiusMap                  map[string]int                                `json:"zoneRadiusMap"`
	fareConfig                     map[string]map[string]*fareConfig             `json:"fareConfig"`
	capConfig                      map[string]map[string]*capConfig              `json:"capConfig"`
	peakHourConfig                 map[string][]*timeRange                       `json:"peakHourConfig.json"`
	returnJourneyOffPeekHourConfig map[string]map[string]map[string][]*timeRange `json:"returnJourneyOffPeekHourConfig"`
}

type fareConfig struct {
	peakFare    int `json:"peakFare"`
	offPeakFare int `json:"offPeakFare"`
}

type timeRange struct {
	fromTime string `json:"fromTime"`
	toTime   string `json:"toTime"`
}

type capConfig struct {
	dailyCap  int `json:"dailyCap"`
	weeklyCap int `json:"weeklyCap"`
}

func (a *appConfig) GetZoneRadius(zoneId string) (int, error) {
	if radius, ok := a.zoneRadiusMap[zoneId]; ok {
		return radius, nil
	} else {
		return -1, errors.New(fmt.Sprintf("invalid zone id : %s", zoneId))
	}
}

func (a *appConfig) isValidZone(zoneId string) bool {
	_, ok := a.zoneRadiusMap[zoneId]
	return ok
}
