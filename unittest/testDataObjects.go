package unittest

type testDataObj struct {
	TestCaseId     string        `json:"testCaseId"`
	Trips          []tripDetails `json:"trips"`
	ExpectedResult int           `json:"expectedResult"`
}

type tripDetails struct {
	DateTimeString string `json:"dateTimeString"`
	FromZone       int    `json:"fromZone"`
	ToZone         int    `json:"toZone"`
}

type peakHoursDataObj struct {
	WeekdayPeakHours map[string]string `json:"weekdayPeakHours"`
	WeekendPeakHours map[string]string `json:"weekendPeakHours"`
}

type zone1ReturnTripOffPeakHoursDataObj struct {
	WeekdayOffPeakHours map[string]string `json:"weekdayOffPeakHours"`
	WeekendOffPeakHours map[string]string `json:"weekendOffPeakHours"`
}

type fareData struct {
	PeakHour    [][]int `json:"peakHour"`
	OffPeakHour [][]int `json:"offPeakHour"`
}

type capData struct {
	Daily  [][]int `json:"daily"`
	Weekly [][]int `json:"weekly"`
}
