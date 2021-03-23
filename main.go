package main

import (
	"flag"
	"fmt"
	"os"
	"tiger-card/cap/daily"
	"tiger-card/cap/weekly"
	"tiger-card/card"
	"tiger-card/config"
	"tiger-card/enum/zones"
	"tiger-card/fare"
	"tiger-card/fare/calculate"
	"tiger-card/logger"
	"tiger-card/peakhr"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

func main() {
	logger.InitLogger()
	var err error

	if err = config.InitializeISTLocation(); err != nil {
		logger.GetLogger().Errorf("Unable to initialize IST location due to error : %+v", err)
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

	logger.GetLogger().Info("Initialization done successfully")

	today := time.Now().In(config.ISTLocation)

	var trips = make([]*trip.Trip, 0, 0)
	tigerCard := card.NewTigerCard()
	trips = append(trips, &trip.Trip{
		CardNumber: tigerCard.GetCardNumber(),
		FromZone:   zones.Z2.Name(),
		ToZone:     zones.Z1.Name(),
		DateTime:   time.Date(today.Year(), today.Month(), today.Day(), 10, 20, 00, 00, today.Location()),
	})

	trips = append(trips, &trip.Trip{
		CardNumber: tigerCard.GetCardNumber(),
		FromZone:   zones.Z1.Name(),
		ToZone:     zones.Z1.Name(),
		DateTime:   time.Date(today.Year(), today.Month(), today.Day(), 10, 45, 00, 00, today.Location()),
	})

	trips = append(trips, &trip.Trip{
		CardNumber: tigerCard.GetCardNumber(),
		FromZone:   zones.Z1.Name(),
		ToZone:     zones.Z1.Name(),
		DateTime:   time.Date(today.Year(), today.Month(), today.Day(), 16, 15, 00, 00, today.Location()),
	})

	trips = append(trips, &trip.Trip{
		CardNumber: tigerCard.GetCardNumber(),
		FromZone:   zones.Z1.Name(),
		ToZone:     zones.Z1.Name(),
		DateTime:   time.Date(today.Year(), today.Month(), today.Day(), 18, 15, 00, 00, today.Location()),
	})

	trips = append(trips, &trip.Trip{
		CardNumber: tigerCard.GetCardNumber(),
		FromZone:   zones.Z1.Name(),
		ToZone:     zones.Z2.Name(),
		DateTime:   time.Date(today.Year(), today.Month(), today.Day(), 19, 00, 00, 00, today.Location()),
	})

	fmt.Println("Total fare is ", calculate.FareCalculator(tigerCard.GetCardNumber(), trips))

}
