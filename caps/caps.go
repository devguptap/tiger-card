package caps

var dailyCap [][]int
var weeklyCap [][]int

func InitCaps(dailyCapMatrix [][]int, weeklyCapMatrix [][]int) {
	dailyCap = dailyCapMatrix
	weeklyCap = weeklyCapMatrix
}

func GetDailyCap(fromZone, toZone int) int {
	return dailyCap[fromZone-1][toZone-1]
}

func GetWeeklyCap(fromZone, toZone int) int {
	return weeklyCap[fromZone-1][toZone-1]
}
