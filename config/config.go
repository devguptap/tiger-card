package config

import (
	"encoding/json"
	"io/ioutil"
	"tiger-card/constants"
	"tiger-card/logger"
	"time"
)

var Config *appConfig
var ISTLocation *time.Location

// InitializeISTLocation initialize the IST time zone (GMT +05:30)
func InitializeISTLocation() error {
	var err error
	ISTLocation, err = time.LoadLocation(constants.IstLocationIANAString)
	return err
}

func InitializeConfig() error {
	var err error
	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile("resources\\config.json"); err == nil {
		Config = new(appConfig)
		if err = json.Unmarshal(fileBytes, Config); err == nil {
			logger.GetLogger().Info("[Config]. Config initialized successfully")
		}
	}

	return err
}
