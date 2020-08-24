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

// Order order
//
// swagger:model order
type Order struct {

	// advanced
	Advanced Advanced `json:"advanced,omitempty"`

	// amount
	Amount Amount `json:"amount,omitempty"`

	// api
	// Required: true
	API API `json:"api"`

	// average price
	AveragePrice AveragePrice `json:"average_price,omitempty"`

	// commission
	Commission Commission `json:"commission,omitempty"`

	// creation timestamp
	// Required: true
	CreationTimestamp Timestamp `json:"creation_timestamp"`

	// direction
	// Required: true
	Direction Direction `json:"direction"`

	// filled amount
	FilledAmount FilledAmount `json:"filled_amount,omitempty"`

	// implv
	Implv Implv `json:"implv,omitempty"`

	// instrument name
	// Required: true
	InstrumentName InstrumentName `json:"instrument_name"`

	// `true` if order was automatically created during liquidation
	// Required: true
	IsLiquidation *bool `json:"is_liquidation"`

	// label
	// Required: true
	Label Label `json:"label"`

	// last update timestamp
	// Required: true
	LastUpdateTimestamp Timestamp `json:"last_update_timestamp"`

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

	// profit loss
	ProfitLoss ProfitLoss `json:"profit_loss,omitempty"`

	// reduce only
	ReduceOnly ReduceOnly `json:"reduce_only,omitempty"`

	// stop price
	StopPrice StopPrice `json:"stop_price,omitempty"`

	// time in force
	// Required: true
	TimeInForce TimeInForce `json:"time_in_force"`

	// trigger
	Trigger Trigger `json:"trigger,omitempty"`

	// triggered
	Triggered Triggered `json:"triggered,omitempty"`

	// usd
	Usd Usd `json:"usd,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvanced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
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

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilledAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImplv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsLiquidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTimestamp(formats); err != nil {
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

	if err := m.validateProfitLoss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReduceOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
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

func (m *Order) validateAdvanced(formats strfmt.Registry) error {

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

func (m *Order) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *Order) validateAPI(formats strfmt.Registry) error {

	if err := m.API.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("api")
		}
		return err
	}

	return nil
}

func (m *Order) validateAveragePrice(formats strfmt.Registry) error {

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

func (m *Order) validateCommission(formats strfmt.Registry) error {

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

func (m *Order) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creation_timestamp")
		}
		return err
	}

	return nil
}

func (m *Order) validateDirection(formats strfmt.Registry) error {

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *Order) validateFilledAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.FilledAmount) { // not required
		return nil
	}

	if err := m.FilledAmount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("filled_amount")
		}
		return err
	}

	return nil
}

func (m *Order) validateImplv(formats strfmt.Registry) error {

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

func (m *Order) validateInstrumentName(formats strfmt.Registry) error {

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

func (m *Order) validateIsLiquidation(formats strfmt.Registry) error {

	if err := validate.Required("is_liquidation", "body", m.IsLiquidation); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateLabel(formats strfmt.Registry) error {

	if err := m.Label.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("label")
		}
		return err
	}

	return nil
}

func (m *Order) validateLastUpdateTimestamp(formats strfmt.Registry) error {

	if err := m.LastUpdateTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("last_update_timestamp")
		}
		return err
	}

	return nil
}

func (m *Order) validateMaxShow(formats strfmt.Registry) error {

	if err := m.MaxShow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_show")
		}
		return err
	}

	return nil
}

func (m *Order) validateOrderID(formats strfmt.Registry) error {

	if err := m.OrderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_id")
		}
		return err
	}

	return nil
}

func (m *Order) validateOrderState(formats strfmt.Registry) error {

	if err := m.OrderState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_state")
		}
		return err
	}

	return nil
}

func (m *Order) validateOrderType(formats strfmt.Registry) error {

	if err := m.OrderType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("order_type")
		}
		return err
	}

	return nil
}

func (m *Order) validatePostOnly(formats strfmt.Registry) error {

	if err := m.PostOnly.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_only")
		}
		return err
	}

	return nil
}

func (m *Order) validatePrice(formats strfmt.Registry) error {

	if err := m.Price.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("price")
		}
		return err
	}

	return nil
}

func (m *Order) validateProfitLoss(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfitLoss) { // not required
		return nil
	}

	if err := m.ProfitLoss.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("profit_loss")
		}
		return err
	}

	return nil
}

func (m *Order) validateReduceOnly(formats strfmt.Registry) error {

	if swag.IsZero(m.ReduceOnly) { // not required
		return nil
	}

	if err := m.ReduceOnly.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reduce_only")
		}
		return err
	}

	return nil
}

func (m *Order) validateStopPrice(formats strfmt.Registry) error {

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

func (m *Order) validateTimeInForce(formats strfmt.Registry) error {

	if err := m.TimeInForce.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time_in_force")
		}
		return err
	}

	return nil
}

func (m *Order) validateTrigger(formats strfmt.Registry) error {

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

func (m *Order) validateTriggered(formats strfmt.Registry) error {

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

func (m *Order) validateUsd(formats strfmt.Registry) error {

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
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
