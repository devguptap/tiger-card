package trip

import (
	"tiger-card/zone"
	"time"
)

type Trip struct {
	FromZone, ToZone zone.Zone
	DateTime         time.Time
}

func NewTrip(from, to int, dateTime time.Time) *Trip {
	return &Trip{
		FromZone: zone.NewZone(from),
		ToZone:   zone.NewZone(to),
		DateTime: dateTime,
	}
}
