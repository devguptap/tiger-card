package unittest

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
	"tiger-card/farecalculator"
	"tiger-card/reqresp"
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
			var trips = make([]*reqresp.TripData, 0, 0)
			var dateTime time.Time
			for _, trip := range testCase.Trips {
				if dateTime, err = time.Parse(testDataDateTimeFormat, trip.DateTimeString); err != nil {
					t.Errorf("Unable to parse date time string : %s due to error : %+v", trip.DateTimeString, err)
					t.FailNow()
				}

				trips = append(trips, &reqresp.TripData{
					DateTime: dateTime,
					FromZone: trip.FromZone,
					ToZone:   trip.ToZone,
				})
			}
			Convey("When run the test case", func() {
				actualResult := farecalculator.CalculateFare(trips)
				Convey(fmt.Sprintf("Then the expected result should be : %v", testCase.ExpectedResult), func() {
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
