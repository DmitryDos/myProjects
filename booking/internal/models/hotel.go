// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Hotel hotel
//
// swagger:model Hotel
type Hotel struct {

	// address
	// Example: Red Square №1
	// Required: true
	Address *string `json:"address"`

	// city
	// Example: Moscow
	// Required: true
	City *string `json:"city"`

	// cost
	Cost int64 `json:"cost,omitempty"`

	// number of stars of hotel
	// Enum: [0,1,2,3,4,5]
	HotelClass int64 `json:"hotel_class,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Radisson
	// Required: true
	Name *string `json:"name"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this hotel
func (m *Hotel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHotelClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hotel) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *Hotel) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

var hotelTypeHotelClassPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hotelTypeHotelClassPropEnum = append(hotelTypeHotelClassPropEnum, v)
	}
}

// prop value enum
func (m *Hotel) validateHotelClassEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, hotelTypeHotelClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Hotel) validateHotelClass(formats strfmt.Registry) error {
	if swag.IsZero(m.HotelClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateHotelClassEnum("hotel_class", "body", m.HotelClass); err != nil {
		return err
	}

	return nil
}

func (m *Hotel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hotel based on context it is used
func (m *Hotel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Hotel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hotel) UnmarshalBinary(b []byte) error {
	var res Hotel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
