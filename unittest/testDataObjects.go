package unittest

type testDataObj struct {
	TestCaseId     string        `json:"testCaseId"`
	Trips          []tripDetails `json:"trips"`
	ExpectedResult int           `json:"expectedResult"`
}

type tripDetails struct {
	DateTimeString string `json:"dateTimeString"`
	FromZone       string `json:"fromZone"`
	ToZone         string `json:"toZone"`
}
