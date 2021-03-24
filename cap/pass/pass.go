package pass

import "time"

// Pass represents the daily and weekly pass.
// If Pass achieved for a tiger-card
// then all ride will be free till the passExpiry
type Pass struct {
	passType   string    `json:"passType"`
	passExpiry time.Time `json:"passExpiry"`
}

// GetPassType returns the type of the pass (daily, weekly)
func (p *Pass) GetPassType() string {
	return p.passType
}

// GetPassExpiry returns the expiry date of the pass
func (p *Pass) GetPassExpiry() time.Time {
	return p.passExpiry
}

// NewPass initialize and return the new pass
func NewPass(passType string, passExpiry time.Time) *Pass {
	return &Pass{
		passType:   passType,
		passExpiry: passExpiry,
	}
}
