package unittest

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"tiger-card/cap/daily"
	"tiger-card/cap/weekly"
	"tiger-card/config"
	"tiger-card/fare"
	"tiger-card/logger"
	"tiger-card/peakhr"
	"tiger-card/zone"
)

func TestMain(m *testing.M) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Test failed with error : %+v\n", r)
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

	if err = fare.InitFare(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize fare due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = daily.InitDailyCap(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize daily cap due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = weekly.InitWeeklyCap(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize weekly cap due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = peakhr.InitPeakHrMap(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize peak hour map due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = peakhr.InitReturnTripOffPeakHrMap(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize return trip off peak hour map due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	if err = zone.InitZones(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize zone due to error : %+v", err)
		os.Exit(int(flag.ExitOnError))
	}

	logger.GetLogger().Info("Starting test cases")
	exitVal := m.Run()
	logger.GetLogger().Infof("Test case execution done. Exit code is : %v", exitVal)
	os.Exit(exitVal)
}
