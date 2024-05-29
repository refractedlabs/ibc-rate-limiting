package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// Validate validates the set of params
func (p *Params) Validate() error {
	for _, admin := range p.Admins {
		if _, err := sdk.AccAddressFromBech32(admin); err != nil {
			return errors.ErrInvalidAddress.Wrapf("invalid admin address: %s", admin)
		}
	}
	return nil
}
