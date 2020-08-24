// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeribitPriceIndexNotification deribit price index notification
//
// swagger:model deribit_price_index_notification
type DeribitPriceIndexNotification struct {

	// index name
	// Required: true
	IndexName *string `json:"index_name"`

	// Current value of Deribit Index
	// Required: true
	Price *float64 `json:"price"`

	// timestamp
	// Required: true
	Timestamp Timestamp `json:"timestamp"`
}

// Validate validates this deribit price index notification
func (m *DeribitPriceIndexNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeribitPriceIndexNotification) validateIndexName(formats strfmt.Registry) error {

	if err := validate.Required("index_name", "body", m.IndexName); err != nil {
		return err
	}

	return nil
}

func (m *DeribitPriceIndexNotification) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *DeribitPriceIndexNotification) validateTimestamp(formats strfmt.Registry) error {

	if err := m.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeribitPriceIndexNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeribitPriceIndexNotification) UnmarshalBinary(b []byte) error {
	var res DeribitPriceIndexNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
