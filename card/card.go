package card

import (
	"tiger-card/interface/icap"
	"tiger-card/trip"
	"time"
)

var TigerCardMap = make(map[string]*TigerCard)

type TigerCard struct {
	number        string       `json:"number"`
	totalFare     int          `json:"totalFare"`
	pass          *Pass        `json:"pass"`
	trips         []*trip.Trip `json:"trips"`
	applicableCap []icap.Cap   `json:"applicableCap"`
}

type Pass struct {
	passType   string    `json:"passType"`
	passExpiry time.Time `json:"passExpiry"`
}

func (p *Pass) GetPassType() string {
	return p.passType
}

func (p *Pass) GetPassExpiry() time.Time {
	return p.passExpiry
}

func NewPass(passType string, passExpiry time.Time) *Pass {
	return &Pass{
		passType:   passType,
		passExpiry: passExpiry,
	}
}

func (c *TigerCard) GetCardNumber() string {
	return c.number
}

func (c *TigerCard) AddToTotalFare(fare int) {
	c.totalFare += fare
}

func (c *TigerCard) GetTotalFare() int {
	return c.totalFare
}

func (c *TigerCard) AddPass(pass *Pass) {
	c.pass = pass
}

func (c *TigerCard) DeletePass() {
	c.pass = nil
}

func (c *TigerCard) AddTrip(t *trip.Trip) {
	if c.trips == nil {
		c.trips = make([]*trip.Trip, 0, 0)
	}
	c.trips = append(c.trips, t)
}

func (c *TigerCard) GetTrips(t *trip.Trip) []*trip.Trip {
	return c.trips
}
