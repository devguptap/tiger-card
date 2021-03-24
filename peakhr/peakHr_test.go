package peakhr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/config"
	"tiger-card/logger"
	"tiger-card/trip"
	"tiger-card/zone"
	"time"
)

func TestPeakHour(t *testing.T) {
	logger.InitLogger()
	err := config.InitializeISTLocation()
	Convey("Given a peak hour config file : peakHourConfig.json", t, func() {
		So(err, ShouldBeNil)
		Convey("When we initialize the peak hour map using config file", func() {
			err = InitPeakHrMap()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And peak hour map should be initialized correctly", func() {
					So(peakHrMap, ShouldNotBeNil)

					So(peakHrMap, ShouldContainKey, time.Monday)
					So(peakHrMap, ShouldContainKey, time.Tuesday)
					So(peakHrMap, ShouldContainKey, time.Wednesday)
					So(peakHrMap, ShouldContainKey, time.Thursday)
					So(peakHrMap, ShouldContainKey, time.Friday)
					So(peakHrMap, ShouldContainKey, time.Saturday)
					So(peakHrMap, ShouldContainKey, time.Sunday)

					Convey("And IsPeakHours(trip *trip.Trip) function should return true for travel time : 22-03-2021 10:00", func() {
						t := &trip.Trip{
							CardNumber: 1122334455,
							FromZone:   "Z1",
							ToZone:     "Z2",
							DateTime:   time.Date(2021, 03, 22, 10, 0, 0, 0, config.ISTLocation),
						}
						So(IsPeakHours(t), ShouldBeTrue)
						Convey("And IsPeakHours(trip *trip.Trip) function should return false for travel time : 22-03-2021 10:31", func() {
							t := &trip.Trip{
								CardNumber: 1122334455,
								FromZone:   "Z1",
								ToZone:     "Z2",
								DateTime:   time.Date(2021, 03, 22, 10, 31, 0, 0, config.ISTLocation),
							}
							So(IsPeakHours(t), ShouldBeFalse)
						})
					})
				})
			})
		})
	})
}

func TestReturnTripOffPeakHour(t *testing.T) {
	logger.InitLogger()
	err := config.InitializeISTLocation()
	Convey("Given a return trip off-peak hour config file : returnTripOffPeakHourConfig.json", t, func() {
		So(err, ShouldBeNil)
		Convey("When we initialize the return trip off-peak hour map using config file", func() {
			err = InitReturnTripOffPeakHrMap()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And return trip off-peak hour map should be initialized correctly", func() {
					So(returnTripOffPeakHrMap, ShouldNotBeNil)

					So(returnTripOffPeakHrMap, ShouldContainKey, zone.Id("Z2"))
					So(returnTripOffPeakHrMap[zone.Id("Z2")], ShouldContainKey, zone.Id("Z1"))
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Monday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Tuesday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Wednesday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Thursday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Friday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Saturday)
					So(returnTripOffPeakHrMap[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, time.Sunday)

					Convey("And isReturnTripOffPeakHr(trip *trip.Trip) function should return true for travel time : 22-03-2021 17:01 and journey from Z2 to Z1", func() {
						t := &trip.Trip{
							CardNumber: 1122334455,
							FromZone:   "Z2",
							ToZone:     "Z1",
							DateTime:   time.Date(2021, 03, 22, 17, 01, 0, 0, config.ISTLocation),
						}
						So(isReturnTripOffPeakHr(t), ShouldBeTrue)

						Convey("And isReturnTripOffPeakHr(trip *trip.Trip) function should return false for 22-03-2021 17:01 but journey from Z1 to Z2", func() {
							t := &trip.Trip{
								CardNumber: 1122334455,
								FromZone:   "Z1",
								ToZone:     "Z2",
								DateTime:   time.Date(2021, 03, 22, 10, 31, 0, 0, config.ISTLocation),
							}
							So(isReturnTripOffPeakHr(t), ShouldBeFalse)
						})
					})
				})
			})
		})
	})
}
