// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// AveragePrice Average fill price of the order
//
// swagger:model average_price
type AveragePrice float64

// Validate validates this average price
func (m AveragePrice) Validate(formats strfmt.Registry) error {
	return nil
}
