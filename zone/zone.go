package zone

import "strconv"

// Zone is a concrete type to represent the zones
type Zone int

func (z Zone) GetId() int {
	return int(z)
}

// NewZone return the Zone object for given zone number/id
func NewZone(id int) Zone {
	return Zone(id)
}

// NewZoneForZoneId return the Zone object for given zone number/id  in string format
func NewZoneForZoneId(zoneId string) Zone {
	var zone Zone
	if id, err := strconv.Atoi(zoneId); err == nil {
		zone = NewZone(id)
	} else {
		panic(err)
	}
	return zone
}
