package util

import "tiger-card/zone"

// GetZoneDistance accepts two zone z1 and z2 and return their absolute difference
func GetZoneDistance(z1, z2 zone.Zone) int {
	if z1.GetId() > z2.GetId() {
		return z1.GetId() - z2.GetId()
	} else {
		return z2.GetId() - z1.GetId()
	}
}
