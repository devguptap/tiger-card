package zones

type zone int

const (
	Z1 zone = iota
	Z2
)

var zoneString = []string{
	"Z1",
	"Z2",
}

func (z zone) Name() string {
	return zoneString[z]
}

func (z zone) String() string {
	return zoneString[z]
}
