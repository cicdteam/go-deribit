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

// OrderType2 Order type, `"all"`, `"limit"`, `"stop_all"`, `"stop_limit"` or `"stop_market"`
//
// swagger:model order_type2
type OrderType2 string

const (

	// OrderType2All captures enum value "all"
	OrderType2All OrderType2 = "all"

	// OrderType2Limit captures enum value "limit"
	OrderType2Limit OrderType2 = "limit"

	// OrderType2StopAll captures enum value "stop_all"
	OrderType2StopAll OrderType2 = "stop_all"

	// OrderType2StopLimit captures enum value "stop_limit"
	OrderType2StopLimit OrderType2 = "stop_limit"

	// OrderType2StopMarket captures enum value "stop_market"
	OrderType2StopMarket OrderType2 = "stop_market"
)

// for schema
var orderType2Enum []interface{}

func init() {
	var res []OrderType2
	if err := json.Unmarshal([]byte(`["all","limit","stop_all","stop_limit","stop_market"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderType2Enum = append(orderType2Enum, v)
	}
}

func (m OrderType2) validateOrderType2Enum(path, location string, value OrderType2) error {
	if err := validate.EnumCase(path, location, value, orderType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this order type2
func (m OrderType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
