// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// CurrencyAmount Amount of funds in given currency
//
// swagger:model currency_amount
type CurrencyAmount float64

// Validate validates this currency amount
func (m CurrencyAmount) Validate(formats strfmt.Registry) error {
	return nil
}
