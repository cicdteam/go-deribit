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

// PrivateGetPositionResponse private get position response
//
// swagger:model private_get_position_response
type PrivateGetPositionResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *Position `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrivateGetPositionResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *Position `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrivateGetPositionResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Result *Position `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this private get position response
func (m *PrivateGetPositionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetPositionResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetPositionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetPositionResponse) UnmarshalBinary(b []byte) error {
	var res PrivateGetPositionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
