package unittest

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
	"tiger-card/config"
	"tiger-card/fare/calculate"
	"tiger-card/trip"
	"time"
)

const (
	testDataDateTimeFormat = "02-01-2006 15:04"
)

func TestTigerCard(t *testing.T) {
	var testData []*testDataObj
	var err error
	if testData, err = getTestData(); err != nil {
		t.Errorf("Unable to parse test data due to error : %+v", err)
		t.FailNow()
	}

	for _, testCase := range testData {
		Convey(fmt.Sprintf("Given a test case with id : %s", testCase.TestCaseId), t, func() {
			var trips = make([]*trip.Trip, 0, 0)
			var dateTime time.Time
			for _, tripTestData := range testCase.Trips {
				if dateTime, err = time.ParseInLocation(testDataDateTimeFormat, tripTestData.DateTimeString, config.ISTLocation); err != nil {
					t.Errorf("Unable to parse date time string : %s due to error : %+v", tripTestData.DateTimeString, err)
					t.FailNow()
				}

				trips = append(trips, trip.NewTrip(tripTestData.FromZone, tripTestData.ToZone, dateTime))
			}
			Convey("When run the test case", func() {
				actualResult := calculate.FareCalculator(trips)
				Convey(fmt.Sprintf("Then the expected result should be : %v", testCase.ExpectedResult), func() {
					t.Logf("Actual Result is : %v", actualResult)
					So(actualResult, ShouldEqual, testCase.ExpectedResult)
				})
			})
		})
	}

}

func getTestData() ([]*testDataObj, error) {
	var err error
	var testData []*testDataObj
	if fileBytes, err := ioutil.ReadFile("./testdata.json"); err == nil {
		err = json.Unmarshal(fileBytes, &testData)
	}
	return testData, err
}
