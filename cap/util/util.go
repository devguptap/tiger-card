package util

import "tiger-card/zone"

func GetZoneDistance(z1, z2 zone.Zone) int {
	if z1.GetId() > z2.GetId() {
		return z1.GetId() - z2.GetId()
	} else {
		return z2.GetId() - z1.GetId()
	}
}
