package fare

var peakHourFare [][]int
var offPeakHourFare [][]int

func InitZoneFare(peakHourFareMatrix, offPeakHourFareMatrix [][]int) {
	peakHourFare = peakHourFareMatrix
	offPeakHourFare = offPeakHourFareMatrix
}

func GetFare(fromZone, toZone int, isPeakHour bool) int {
	if isPeakHour {
		return peakHourFare[fromZone-1][toZone-1]
	} else {
		return offPeakHourFare[fromZone-1][toZone-1]
	}
}
