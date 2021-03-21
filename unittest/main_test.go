package unittest

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"runtime/debug"
	"testing"
	"tiger-card/config"
	"tiger-card/fare"
	"tiger-card/logger"
	"tiger-card/pass/daily"
	"tiger-card/pass/weekly"
	"tiger-card/peakhr"
)

func TestMain(m *testing.M) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetLogger().Errorf("Test failed with error : %+v", r)
			debug.PrintStack()
			os.Exit(int(flag.ExitOnError))
		}
	}()
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

	logger.GetLogger().Info("Starting test cases")
	exitVal := m.Run()
	logger.GetLogger().Infof("Test case execution done. Exit code is : %v", exitVal)
	os.Exit(exitVal)
}

func initPeakHourTestData() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./peekHourConfig.json"); err == nil {
		peakhr.InitializePeakHour(fileBytes)
	}
	return err
}

func initCapData() error {
	var fileBytes []byte
	var err error
	if fileBytes, err = ioutil.ReadFile("./capData.json"); err == nil {
		dataObj := &capData{}
		if err = json.Unmarshal(fileBytes, dataObj); err == nil {
			daily.Init(dataObj.Daily)
			weekly.Init(dataObj.Weekly)
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
			fare.Init(dataObj.PeakHour, dataObj.OffPeakHour)
		}
	}
	return err
}
