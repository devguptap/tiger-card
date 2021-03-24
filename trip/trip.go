package trip

import (
	"time"
)

// Trip represent a trip object containing tiger-card number, from zone, to zone and time of travel.
type Trip struct {
	CardNumber       int
	FromZone, ToZone string
	DateTime         time.Time
}

// NewTrip accepts the from,to zones and dateTime and initialize and return the trip object
func NewTrip(cardNumber int, fromZone, toZone string, dateTime time.Time) *Trip {
	return &Trip{
		CardNumber: cardNumber,
		FromZone:   fromZone,
		ToZone:     toZone,
		DateTime:   dateTime,
	}
}
