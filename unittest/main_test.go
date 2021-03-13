package unittest

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"
	"tiger-card/caps"
	"tiger-card/config"
	"tiger-card/fare"
	"tiger-card/logger"
	"tiger-card/peakhours"
)

func TestMain(m *testing.M) {
	logger.InitLogger()
	var err error
	if err = config.InitializeISTLocation(); err != nil {
		logger.GetLogger().Errorf("Unable to parse IST timezone due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = initFareData(); err != nil {
		logger.GetLogger().Errorf("Unable to parse fare data file due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = initCapData(); err != nil {
		logger.GetLogger().Errorf("Unable to parse fare cap data file due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = initPeakHourTestData(); err != nil {
		logger.GetLogger().Errorf("Unable to parse peak hour data file due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = initZone1ReturnTripOffPeakHours(); err != nil {
		logger.GetLogger().Errorf("Unable to parse zone 1 return trio off peak hour data file due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	logger.GetLogger().Info("Starting test cases")
	exitVal := m.Run()
	logger.GetLogger().Infof("Test case execution done. Exit code is : %v", exitVal)
	os.Exit(exitVal)
}

func initPeakHourTestData() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./peakHourData.json"); err == nil {
		dataObj := &peakHoursDataObj{}
		if err = json.Unmarshal(fileBytes, dataObj); err == nil {
			err = peakhours.InitPeakHours(dataObj.WeekdayPeakHours, dataObj.WeekendPeakHours)
		}
	}
	return err
}

func initZone1ReturnTripOffPeakHours() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./zone1ReturnTripOffPeakHours.json"); err == nil {
		dataObj := &zone1ReturnTripOffPeakHoursDataObj{}
		if err = json.Unmarshal(fileBytes, dataObj); err == nil {
			err = peakhours.InitZone1ReturnJourneyOffPeeHours(dataObj.WeekdayOffPeakHours, dataObj.WeekendOffPeakHours)
		}
	}
	return err
}

func initCapData() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./capData.json"); err == nil {
		dataObj := &capData{}
		if err = json.Unmarshal(fileBytes, dataObj); err == nil {
			caps.InitCaps(dataObj.Daily, dataObj.Weekly)
		}
	}
	return err
}

func initFareData() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./fareData.json"); err == nil {
		dataObj := &fareData{}
		if err = json.Unmarshal(fileBytes, dataObj); err == nil {
			fare.InitZoneFare(dataObj.PeakHour, dataObj.OffPeakHour)
		}
	}
	return err
}
