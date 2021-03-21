package cache

import "time"

const (
	dailyTotalSuffix  = "_daily_total"
	weeklyTotalSuffix = "_weekly_total"
	//monthlyTotalSuffix = "monthly_total"
)

var cache = make(map[string]*Value)

type Value struct {
	amount int
	expiry time.Time
}

func GetDailyTotal(cardNumber string, timestamp time.Time) int {
	if value, ok := cache[cardNumber+dailyTotalSuffix]; ok {
		if value.expiry.After(timestamp) {
			return value.amount
		}
	}

	return 0
}

func UpdateDailyTotal(cardNumber string, fare int, timestamp time.Time) {
	if value, ok := cache[cardNumber+dailyTotalSuffix]; ok {
		if value.expiry.After(timestamp) {
			value.amount += fare
		} else {
			value = new(Value)
			value.amount = fare
			value.expiry = time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day()+1, 0, 0, 0, 0, timestamp.Location())
			cache[cardNumber+dailyTotalSuffix] = value
		}
	}
}

func GetWeeklyTotal(cardNumber string, timestamp time.Time) int {
	if value, ok := cache[cardNumber+weeklyTotalSuffix]; ok {
		if value.expiry.After(timestamp) {
			return value.amount
		}
	}

	return 0
}

func UpdateWeeklyTotal(cardNumber string, fare int, timestamp time.Time) {
	if value, ok := cache[cardNumber+weeklyTotalSuffix]; ok {
		if value.expiry.After(timestamp) {
			value.amount += fare
		} else {
			value = new(Value)
			value.amount = fare
			value.expiry = getEndOfTheWeekForDate(timestamp)
			cache[cardNumber+weeklyTotalSuffix] = value
		}
	}
}

func getEndOfTheWeekForDate(t time.Time) time.Time {
	dayOfTheWeek := int(t.Weekday())
	dayToAdd := (7-dayOfTheWeek)%7 + 1
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, dayToAdd)
}
