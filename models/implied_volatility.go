// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// ImpliedVolatility Value of the volatility of the underlying instrument
//
// swagger:model implied_volatility
type ImpliedVolatility float64

// Validate validates this implied volatility
func (m ImpliedVolatility) Validate(formats strfmt.Registry) error {
	return nil
}
