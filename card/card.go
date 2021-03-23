package card

import (
	"math/rand"
	"tiger-card/cap/pass"
	"tiger-card/interface/icap"
	"tiger-card/trip"
	"time"
)

var TigerCardMap = make(map[int]*TigerCard)

type TigerCard struct {
	number        int          `json:"number"`
	totalFare     int          `json:"totalFare"`
	pass          *pass.Pass   `json:"pass"`
	trips         []*trip.Trip `json:"trips"`
	applicableCap []icap.Cap   `json:"applicableCap"`
}

func (c *TigerCard) GetCardNumber() int {
	return c.number
}

func (c *TigerCard) AddToTotalFare(fare int) {
	c.totalFare += fare
}

func (c *TigerCard) GetTotalFare() int {
	return c.totalFare
}

func (c *TigerCard) GetApplicableCap() []icap.Cap {
	return c.applicableCap
}

func (c *TigerCard) AddPass(pass *pass.Pass) {
	c.pass = pass
}

func (c *TigerCard) deletePass() {
	c.pass = nil
}

func (c *TigerCard) AddTrip(t *trip.Trip) {
	if c.trips == nil {
		c.trips = make([]*trip.Trip, 0, 0)
	}
	c.trips = append(c.trips, t)
}

func (c *TigerCard) GetTrips() []*trip.Trip {
	return c.trips
}

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

func NewTigerCard() *TigerCard {
	card := &TigerCard{
		number:        randomInt(100000, 999999999),
		totalFare:     0,
		pass:          nil,
		trips:         nil,
		applicableCap: icap.GetCaps(),
	}
	//	resetAllCaps(card.applicableCap, time.Unix(0, 0))
	TigerCardMap[card.number] = card
	return card

}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// resetAllCaps reset all the caps
func resetAllCaps(caps []icap.Cap, t *trip.Trip) {
	for _, c := range caps {
		c.Reset(t)
	}
}
