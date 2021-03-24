package zones

// zone represents a zone constant
type zone int

const (
	Z1 zone = iota
	Z2
)

// zoneString contains the string representation of zone for that zone constant
var zoneString = []string{
	"Z1",
	"Z2",
}

// Name returns the zone string for zone constant
func (z zone) Name() string {
	return zoneString[z]
}

// String returns the zone string for zone constant
func (z zone) String() string {
	return zoneString[z]
}
