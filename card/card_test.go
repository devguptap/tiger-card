package card

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"tiger-card/logger"
)

func TestCard(t *testing.T) {
	Convey("Given a tiger-card map", t, func() {
		logger.InitLogger()
		Convey("When we initialize a new tiger-card", func() {
			tigerCard := NewTigerCard()
			Convey("Then received tiger-card should not be nil", func() {
				So(tigerCard, ShouldNotBeNil)
				Convey(fmt.Sprintf("And this tiger-card should be present in TigerCardMap against tiger-card number : %d", tigerCard.number), func() {
					So(TigerCardMap[tigerCard.number], ShouldResemble, tigerCard)
					Convey(fmt.Sprintf("And GetCardNumber() should return card number : %d", tigerCard.number), func() {
						cardNum := tigerCard.GetCardNumber()
						So(cardNum, ShouldEqual, tigerCard.number)
						Convey(fmt.Sprintf("And AddToTotalFare(3) should add 3 to the total fare"), func() {
							tigerCard.AddToTotalFare(3)
							So(tigerCard.totalFare, ShouldEqual, 3)
							tigerCard.AddToTotalFare(7)
							So(tigerCard.totalFare, ShouldEqual, 10)

							Convey("And GetTotalFare() should return the correct fare", func() {
								So(tigerCard.GetTotalFare(), ShouldEqual, 10)
							})
						})
					})
				})
			})
		})
	})
}
