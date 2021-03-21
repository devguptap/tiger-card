package trip

import (
	"tiger-card/zone"
	"time"
)

type Trip struct {
	CardNumber       string
	FromZone, ToZone zone.Zone
	DateTime         time.Time
}

// NewTrip accepts the from,to zones and dateTime and initialize and return the trip object
func NewTrip(cardNumber string, from, to int, dateTime time.Time) *Trip {
	return &Trip{
		CardNumber: cardNumber,
		FromZone:   zone.NewZone(from),
		ToZone:     zone.NewZone(to),
		DateTime:   dateTime,
	}
}
