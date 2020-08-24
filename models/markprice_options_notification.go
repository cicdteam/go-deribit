// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarkpriceOptionsNotification markprice options notification
//
// swagger:model markprice_options_notification
type MarkpriceOptionsNotification []*MarkpriceOptionsNotificationItems0

// Validate validates this markprice options notification
func (m MarkpriceOptionsNotification) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarkpriceOptionsNotificationItems0 markprice options notification items0
//
// swagger:model MarkpriceOptionsNotificationItems0
type MarkpriceOptionsNotificationItems0 struct {

	// instrument name
	InstrumentName InstrumentName `json:"instrument_name,omitempty"`

	// iv
	Iv ImpliedVolatility `json:"iv,omitempty"`

	// mark price
	MarkPrice MarkPrice `json:"mark_price,omitempty"`
}

// Validate validates this markprice options notification items0
func (m *MarkpriceOptionsNotificationItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarkpriceOptionsNotificationItems0) validateInstrumentName(formats strfmt.Registry) error {

	if swag.IsZero(m.InstrumentName) { // not required
		return nil
	}

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

func (m *MarkpriceOptionsNotificationItems0) validateIv(formats strfmt.Registry) error {

	if swag.IsZero(m.Iv) { // not required
		return nil
	}

	if err := m.Iv.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("iv")
		}
		return err
	}

	return nil
}

func (m *MarkpriceOptionsNotificationItems0) validateMarkPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.MarkPrice) { // not required
		return nil
	}

	if err := m.MarkPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mark_price")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarkpriceOptionsNotificationItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarkpriceOptionsNotificationItems0) UnmarshalBinary(b []byte) error {
	var res MarkpriceOptionsNotificationItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
