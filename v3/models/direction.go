// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Direction direction, `buy` or `sell`
//
// swagger:model direction
type Direction string

const (

	// DirectionBuy captures enum value "buy"
	DirectionBuy Direction = "buy"

	// DirectionSell captures enum value "sell"
	DirectionSell Direction = "sell"
)

// for schema
var directionEnum []interface{}

func init() {
	var res []Direction
	if err := json.Unmarshal([]byte(`["buy","sell"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directionEnum = append(directionEnum, v)
	}
}

func (m Direction) validateDirectionEnum(path, location string, value Direction) error {
	if err := validate.EnumCase(path, location, value, directionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this direction
func (m Direction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
