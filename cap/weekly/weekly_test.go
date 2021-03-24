package weekly

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/logger"
	"tiger-card/zone"
)

func TestWeekly(t *testing.T) {
	Convey("Given a weekly cap config file : weeklyCap.json", t, func() {
		logger.InitLogger()
		Convey("When we initialize the weekly cap using config file", func() {
			err := InitWeeklyCap()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And weekly cap map should be initialized correctly", func() {
					So(weeklyCapMap, ShouldNotBeNil)
					So(weeklyCapMap, ShouldContainKey, zone.Id("Z1"))
					So(weeklyCapMap, ShouldContainKey, zone.Id("Z2"))
					So(weeklyCapMap[zone.Id("Z1")], ShouldContainKey, zone.Id("Z1"))
					So(weeklyCapMap[zone.Id("Z1")], ShouldContainKey, zone.Id("Z2"))
					So(weeklyCapMap[zone.Id("Z2")], ShouldContainKey, zone.Id("Z1"))
					So(weeklyCapMap[zone.Id("Z2")], ShouldContainKey, zone.Id("Z2"))
					So(weeklyCapMap[zone.Id("Z1")][zone.Id("Z1")], ShouldEqual, 500)
					So(weeklyCapMap[zone.Id("Z1")][zone.Id("Z2")], ShouldEqual, 600)
					So(weeklyCapMap[zone.Id("Z2")][zone.Id("Z1")], ShouldEqual, 600)
					So(weeklyCapMap[zone.Id("Z2")][zone.Id("Z2")], ShouldEqual, 400)
				})
			})
		})
	})
}
