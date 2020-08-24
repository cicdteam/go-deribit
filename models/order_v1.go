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

// OrderV1 order v1
//
// swagger:model order_v1
type OrderV1 struct {

	// advanced
	Advanced Advanced `json:"advanced,omitempty"`

	// api
	// Required: true
	API API `json:"api"`

	// average price
	AveragePrice AveragePrice `json:"average_price,omitempty"`

	// commission
	Commission Commission `json:"commission,omitempty"`

	// created
	// Required: true
	Created Timestamp `json:"created"`

	// direction
	// Required: true
	Direction Direction `json:"direction"`

	// filled quantity
	FilledQuantity FilledQuantity `json:"filled_quantity,omitempty"`

	// implv
	Implv Implv `json:"implv,omitempty"`

	// instrument
	// Required: true
	Instrument *string `json:"instrument"`

	// label
	// Required: true
	Label Label `json:"label"`

	// last update
	// Required: true
	LastUpdate Timestamp `json:"last_update"`

	// max show
	// Required: true
	MaxShow MaxShow `json:"max_show"`

	// order id
	// Required: true
	OrderID OrderID `json:"order_id"`

	// order state
	// Required: true
	OrderState OrderState `json:"order_state"`

	// order type
	// Required: true
	OrderType OrderType `json:"order_type"`

	// post only
	// Required: true
	PostOnly PostOnly `json:"post_only"`

	// price
	// Required: true
	Price Price `json:"price"`

	// quantity
	Quantity Quantity `json:"quantity,omitempty"`

	// stop price
	StopPrice StopPrice `json:"stop_price,omitempty"`

	// trigger
	Trigger Trigger `json:"trigger,omitempty"`

	// triggered
	Triggered Triggered `json:"triggered,omitempty"`

	// usd
	Usd Usd `json:"usd,omitempty"`
}

// Validate validates this order v1
func (m *OrderV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvanced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAveragePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilledQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImplv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxShow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderV1) validateAdvanced(formats strfmt.Registry) error {

	if swag.IsZero(m.Advanced) { // not required
		return nil
	}

	if err := m.Advanced.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("advanced")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateAPI(formats strfmt.Registry) error {

	if err := m.API.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateAveragePrice(formats strfmt.Registry) error {

	if swag.IsZero(m.AveragePrice) { // not required
		return nil
	}

	if err := m.AveragePrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("average_price")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateCommission(formats strfmt.Registry) error {

	if swag.IsZero(m.Commission) { // not required
		return nil
	}

	if err := m.Commission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("commission")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateCreated(formats strfmt.Registry) error {

	if err := m.Created.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("created")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateDirection(formats strfmt.Registry) error {

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateFilledQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.FilledQuantity) { // not required
		return nil
	}

	if err := m.FilledQuantity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("filled_quantity")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateImplv(formats strfmt.Registry) error {

	if swag.IsZero(m.Implv) { // not required
		return nil
	}

	if err := m.Implv.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("implv")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateInstrument(formats strfmt.Registry) error {

	if err := validate.Required("instrument", "body", m.Instrument); err != nil {
		return err
	}

	return nil
}

func (m *OrderV1) validateLabel(formats strfmt.Registry) error {

	if err := m.Label.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("label")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateLastUpdate(formats strfmt.Registry) error {

	if err := m.LastUpdate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("last_update")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateMaxShow(formats strfmt.Registry) error {

	if err := m.MaxShow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_show")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateOrderID(formats strfmt.Registry) error {

	if err := m.OrderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_id")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateOrderState(formats strfmt.Registry) error {

	if err := m.OrderState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_state")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateOrderType(formats strfmt.Registry) error {

	if err := m.OrderType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_type")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validatePostOnly(formats strfmt.Registry) error {

	if err := m.PostOnly.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_only")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validatePrice(formats strfmt.Registry) error {

	if err := m.Price.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("price")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	if err := m.Quantity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quantity")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateStopPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.StopPrice) { // not required
		return nil
	}

	if err := m.StopPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("stop_price")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if err := m.Trigger.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trigger")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateTriggered(formats strfmt.Registry) error {

	if swag.IsZero(m.Triggered) { // not required
		return nil
	}

	if err := m.Triggered.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("triggered")
		}
		return err
	}

	return nil
}

func (m *OrderV1) validateUsd(formats strfmt.Registry) error {

	if swag.IsZero(m.Usd) { // not required
		return nil
	}

	if err := m.Usd.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("usd")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderV1) UnmarshalBinary(b []byte) error {
	var res OrderV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
