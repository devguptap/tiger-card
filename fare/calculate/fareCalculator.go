package calculate

import (
	"errors"
	"fmt"
	"tiger-card/card"
	"tiger-card/fare"
	"tiger-card/logger"
	"tiger-card/trip"
)

// FareCalculator accepts the list of trips and return the total fare for all the trip
func FareCalculator(cardNumber int, trips []*trip.Trip) int {
	var totalFare int
	var tigerCard *card.TigerCard
	var err error
	for _, t := range trips {
		if tigerCard, err = processAndUpdateTigerCard(cardNumber, t); err != nil {
			logger.GetLogger().Error("Unable to process trip due to error : %+v", err)
			break
		}
	}

	if tigerCard != nil {
		totalFare = tigerCard.GetTotalFare()
	}
	return totalFare
}

func processAndUpdateTigerCard(cardNumber int, t *trip.Trip) (*card.TigerCard, error) {
	var tigerCard *card.TigerCard
	var err error
	var ok bool
	if tigerCard, ok = card.TigerCardMap[cardNumber]; ok {
		tigerCard.AddTrip(t)
		if ok := tigerCard.IsPassAvailable(t); ok {
			return tigerCard, nil
		} else {

			actualFare := fare.GetFare(t)
			for _, c := range tigerCard.GetApplicableCap() {
				c.UpdateCap(t)
				if isCapReached := c.IsCapLimitReached(t, actualFare); isCapReached {
					tigerCard.AddPass(c.GetPass(t))
					actualFare = c.GetCappedFare(actualFare)
					break
				}
			}

			for _, c := range tigerCard.GetApplicableCap() {
				c.UpdateTotalFare(actualFare)
			}

			tigerCard.AddToTotalFare(actualFare)
		}

	} else {
		err = errors.New(fmt.Sprintf("Invalid tiger card number : %d", cardNumber))
	}
	return tigerCard, err
}
