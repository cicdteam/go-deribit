// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateGetDepositsResponse private get deposits response
//
// swagger:model private_get_deposits_response
type PrivateGetDepositsResponse struct {

	// result
	// Required: true
	Result *PrivateGetDepositsResponseResult `json:"result"`
}

// Validate validates this private get deposits response
func (m *PrivateGetDepositsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetDepositsResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PrivateGetDepositsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetDepositsResponse) UnmarshalBinary(b []byte) error {
	var res PrivateGetDepositsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateGetDepositsResponseResult private get deposits response result
//
// swagger:model PrivateGetDepositsResponseResult
type PrivateGetDepositsResponseResult struct {

	// count
	// Required: true
	Count ResultCount `json:"count"`

	// data
	// Required: true
	Data []*Deposit `json:"data"`
}

// Validate validates this private get deposits response result
func (m *PrivateGetDepositsResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetDepositsResponseResult) validateCount(formats strfmt.Registry) error {

	if err := m.Count.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result" + "." + "count")
		}
		return err
	}

	return nil
}

func (m *PrivateGetDepositsResponseResult) validateData(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetDepositsResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetDepositsResponseResult) UnmarshalBinary(b []byte) error {
	var res PrivateGetDepositsResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
