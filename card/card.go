package card

import (
	"math/rand"
	"tiger-card/cap/pass"
	"tiger-card/interface/icap"
	"tiger-card/trip"
	"time"
)

// TigerCardMap contains all the tiger card present in the system. This can be saved in DB later.
var TigerCardMap = make(map[int]*TigerCard)

// TigerCard represents a tiger card and its properties.
type TigerCard struct {
	number        int          `json:"number"`        // Card number
	totalFare     int          `json:"totalFare"`     // total fare for all the journey
	pass          *pass.Pass   `json:"pass"`          // Daily / weekly pass
	trips         []*trip.Trip `json:"trips"`         // All the trips done using that pass
	applicableCap []icap.Cap   `json:"applicableCap"` // applicable cap for a card
}

// GetCardNumber return the card number
func (c *TigerCard) GetCardNumber() int {
	return c.number
}

// AddToTotalFare adds the fare to the totalFare for the tiger-card
func (c *TigerCard) AddToTotalFare(fare int) {
	c.totalFare += fare
}

// GetTotalFare returns the totalFare for the tiger-card
func (c *TigerCard) GetTotalFare() int {
	return c.totalFare
}

// GetApplicableCap returns all cap applicable for the tiger-card
func (c *TigerCard) GetApplicableCap() []icap.Cap {
	return c.applicableCap
}

// AddPass adds a pass to the tiger-card
func (c *TigerCard) AddPass(pass *pass.Pass) {
	c.pass = pass
}

// deletePass deleted the pass from teh card
func (c *TigerCard) deletePass() {
	c.pass = nil
}

// AddTrip add the current trip to the card trip list
func (c *TigerCard) AddTrip(t *trip.Trip) {
	if c.trips == nil {
		c.trips = make([]*trip.Trip, 0, 0)
	}
	c.trips = append(c.trips, t)
}

// GetTrips returns all the trip done on the tiger-card
func (c *TigerCard) GetTrips() []*trip.Trip {
	return c.trips
}

// IsPassAvailable checks if a Pass is available for the tiger-card
func (c *TigerCard) IsPassAvailable(t *trip.Trip) bool {
	if c.pass != nil {
		if c.pass.GetPassExpiry().After(t.DateTime) {
			return true
		} else {
			c.deletePass()
			return false
		}
	}
	return false
}

// NewTigerCard initializes and returns a new card.
func NewTigerCard() *TigerCard {
	card := &TigerCard{
		number:        randomInt(100000, 999999999),
		totalFare:     0,
		pass:          nil,
		trips:         nil,
		applicableCap: icap.GetCaps(),
	}

	TigerCardMap[card.number] = card
	return card

}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
