package zone

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/logger"
)

func TestZone(t *testing.T) {
	logger.InitLogger()
	Convey("Given a zone config file : zones.json", t, func() {

		Convey("When we initialize the zone map using config file", func() {
			err := InitZones()
			Convey("Then error should be nil", func() {
				So(err, ShouldBeNil)
				Convey("And zone map should be initialized correctly", func() {
					So(zoneRadiusMap, ShouldNotBeNil)

					So(zoneRadiusMap, ShouldContainKey, Id("Z1"))
					So(zoneRadiusMap, ShouldContainKey, Id("Z2"))

					So(zoneRadiusMap[Id("Z1")], ShouldEqual, 2)
					So(zoneRadiusMap[Id("Z2")], ShouldEqual, 5)
					Convey("And GetZoneDistance(fromZone, toZone string) function should return 3 for zone Z1 and Z2 given in any order", func() {
						So(GetZoneDistance("Z1", "Z2"), ShouldEqual, 3)
						So(GetZoneDistance("Z2", "Z1"), ShouldEqual, 3)
					})
				})
			})
		})
	})
}
