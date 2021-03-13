package reqresp

import "time"

type TripData struct {
	DateTime time.Time
	FromZone int
	ToZone   int
}
