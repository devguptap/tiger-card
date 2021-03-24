package daily

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/logger"
	"tiger-card/zone"
)

func TestDaily(t *testing.T) {
	Convey("Given a daily cap config file : dailyCap.json", t, func() {
		logger.InitLogger()
		Convey("When we initialize the daily cap using config file", func() {
			err := InitDailyCap()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And daily cap map should be initialized correctly", func() {
					So(dailyCapMap, ShouldNotBeNil)
					So(dailyCapMap, ShouldContainKey, zone.Id("Z1"))
					So(dailyCapMap, ShouldContainKey, zone.Id("Z2"))
					So(dailyCapMap[zone.Id("Z1")], ShouldContainKey, zone.Id("Z1"))
					So(dailyCapMap[zone.Id("Z1")], ShouldContainKey, zone.Id("Z2"))
					So(dailyCapMap[zone.Id("Z2")], ShouldContainKey, zone.Id("Z1"))
					So(dailyCapMap[zone.Id("Z2")], ShouldContainKey, zone.Id("Z2"))
					So(dailyCapMap[zone.Id("Z1")][zone.Id("Z1")], ShouldEqual, 100)
					So(dailyCapMap[zone.Id("Z1")][zone.Id("Z2")], ShouldEqual, 120)
					So(dailyCapMap[zone.Id("Z2")][zone.Id("Z1")], ShouldEqual, 120)
					So(dailyCapMap[zone.Id("Z2")][zone.Id("Z2")], ShouldEqual, 80)
				})
			})
		})
	})
}
