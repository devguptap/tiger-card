package config

import (
	"tiger-card/constants"
	"time"
)

var ISTLocation *time.Location

// InitializeISTLocation initialize the IST time zone (GMT +05:30)
func InitializeISTLocation() error {
	var err error
	ISTLocation, err = time.LoadLocation(constants.IstLocationIANAString)
	return err
}
