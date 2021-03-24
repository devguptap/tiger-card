package fare

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/logger"
	"tiger-card/zone"
)

func TestFare(t *testing.T) {
	Convey("Given a fare config file : fare.json", t, func() {
		logger.InitLogger()
		Convey("When we initialize the fare using config file", func() {
			err := InitFare()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And fare map should be initialized correctly", func() {
					So(fares, ShouldNotBeNil)

					So(fares, ShouldContainKey, zone.Id("Z1"))
					So(fares, ShouldContainKey, zone.Id("Z2"))

					So(fares[zone.Id("Z1")], ShouldContainKey, zone.Id("Z1"))
					So(fares[zone.Id("Z1")], ShouldContainKey, zone.Id("Z2"))
					So(fares[zone.Id("Z2")], ShouldContainKey, zone.Id("Z1"))
					So(fares[zone.Id("Z2")], ShouldContainKey, zone.Id("Z2"))

					So(fares[zone.Id("Z1")][zone.Id("Z1")], ShouldContainKey, peakFare)
					So(fares[zone.Id("Z1")][zone.Id("Z1")], ShouldContainKey, offPeakFare)
					So(fares[zone.Id("Z1")][zone.Id("Z2")], ShouldContainKey, peakFare)
					So(fares[zone.Id("Z1")][zone.Id("Z2")], ShouldContainKey, offPeakFare)
					So(fares[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, peakFare)
					So(fares[zone.Id("Z2")][zone.Id("Z1")], ShouldContainKey, offPeakFare)
					So(fares[zone.Id("Z2")][zone.Id("Z2")], ShouldContainKey, peakFare)
					So(fares[zone.Id("Z2")][zone.Id("Z2")], ShouldContainKey, offPeakFare)

					So(fares[zone.Id("Z1")][zone.Id("Z1")][peakFare], ShouldEqual, 30)
					So(fares[zone.Id("Z1")][zone.Id("Z1")][offPeakFare], ShouldEqual, 25)
					So(fares[zone.Id("Z1")][zone.Id("Z2")][peakFare], ShouldEqual, 35)
					So(fares[zone.Id("Z1")][zone.Id("Z2")][offPeakFare], ShouldEqual, 30)
					So(fares[zone.Id("Z2")][zone.Id("Z1")][peakFare], ShouldEqual, 35)
					So(fares[zone.Id("Z2")][zone.Id("Z1")][offPeakFare], ShouldEqual, 30)
					So(fares[zone.Id("Z2")][zone.Id("Z2")][peakFare], ShouldEqual, 25)
					So(fares[zone.Id("Z2")][zone.Id("Z2")][offPeakFare], ShouldEqual, 20)
				})
			})
		})
	})
}
