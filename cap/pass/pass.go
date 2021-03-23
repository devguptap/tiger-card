package pass

import "time"

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
