package config

import (
	"tiger-card/constants"
	"time"
)

var ISTLocation *time.Location

func InitializeISTLocation() error {
	var err error
	ISTLocation, err = time.LoadLocation(constants.IstLocationIANAString)
	return err
}
