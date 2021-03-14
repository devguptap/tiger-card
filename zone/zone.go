package zone

import "strconv"

type Zone int

func (z Zone) GetId() int {
	return int(z)
}

func NewZone(id int) Zone {
	return Zone(id)
}

func NewZoneForZoneId(zoneId string) Zone {
	var zone Zone
	if id, err := strconv.Atoi(zoneId); err == nil {
		zone = NewZone(id)
	} else {
		panic(err)
	}
	return zone
}
