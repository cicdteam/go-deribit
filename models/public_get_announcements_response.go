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

// PublicGetAnnouncementsResponse public get announcements response
//
// swagger:model public_get_announcements_response
type PublicGetAnnouncementsResponse struct {
	BaseMessage

	// result
	// Required: true
	Result []*PublicGetAnnouncementsResponseResultItems0 `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PublicGetAnnouncementsResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result []*PublicGetAnnouncementsResponseResultItems0 `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PublicGetAnnouncementsResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Result []*PublicGetAnnouncementsResponseResultItems0 `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this public get announcements response
func (m *PublicGetAnnouncementsResponse) Validate(formats strfmt.Registry) error {
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

func (m *PublicGetAnnouncementsResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicGetAnnouncementsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetAnnouncementsResponse) UnmarshalBinary(b []byte) error {
	var res PublicGetAnnouncementsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicGetAnnouncementsResponseResultItems0 public get announcements response result items0
//
// swagger:model PublicGetAnnouncementsResponseResultItems0
type PublicGetAnnouncementsResponseResultItems0 struct {

	// The HTML body of the announcement
	// Required: true
	Body *string `json:"body"`

	// A unique identifier for the announcement
	// Required: true
	ID *float64 `json:"id"`

	// Whether the announcement is marked as important
	// Required: true
	Important *bool `json:"important"`

	// The timestamp in ms at which the announcement was published
	// Required: true
	PublicationTime *int64 `json:"publication_time"`

	// The title of the announcement
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this public get announcements response result items0
func (m *PublicGetAnnouncementsResponseResultItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicGetAnnouncementsResponseResultItems0) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetAnnouncementsResponseResultItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetAnnouncementsResponseResultItems0) validateImportant(formats strfmt.Registry) error {

	if err := validate.Required("important", "body", m.Important); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetAnnouncementsResponseResultItems0) validatePublicationTime(formats strfmt.Registry) error {

	if err := validate.Required("publication_time", "body", m.PublicationTime); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetAnnouncementsResponseResultItems0) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicGetAnnouncementsResponseResultItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetAnnouncementsResponseResultItems0) UnmarshalBinary(b []byte) error {
	var res PublicGetAnnouncementsResponseResultItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}