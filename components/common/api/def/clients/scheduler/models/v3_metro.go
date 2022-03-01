// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3Metro v3 metro
//
// swagger:model v3Metro
type V3Metro struct {

	// City
	//
	// City of the location
	// Read Only: true
	City string `json:"city,omitempty"`

	// Country
	//
	// country of the location
	// Read Only: true
	Country string `json:"country,omitempty"`

	// CountryCode
	//
	// CountryCode of the location
	// Read Only: true
	CountryCode string `json:"countryCode,omitempty"`

	// ID of Location
	//
	// ID Location of the cluster
	ID string `json:"id,omitempty"`

	// Latitude
	//
	// Latitude of the location
	// Read Only: true
	Latitude string `json:"latitude,omitempty"`

	// Locale
	//
	// locale of the location
	// Read Only: true
	Locale string `json:"locale,omitempty"`

	// Longitude
	//
	// Longitude of the location
	// Read Only: true
	Longitude string `json:"longitude,omitempty"`

	// Location
	//
	// Location of the cluster
	Name string `json:"name,omitempty"`

	// State
	//
	// State of the location
	// Read Only: true
	State string `json:"state,omitempty"`

	// StateCode
	//
	// StateCode of the location
	// Read Only: true
	StateCode string `json:"stateCode,omitempty"`
}

// Validate validates this v3 metro
func (m *V3Metro) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this v3 metro based on the context it is used
func (m *V3Metro) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountryCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatitude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLongitude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Metro) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "city", "body", string(m.City)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateCountryCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "countryCode", "body", string(m.CountryCode)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateLatitude(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "latitude", "body", string(m.Latitude)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateLocale(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "locale", "body", string(m.Locale)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateLongitude(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "longitude", "body", string(m.Longitude)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *V3Metro) contextValidateStateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "stateCode", "body", string(m.StateCode)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Metro) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Metro) UnmarshalBinary(b []byte) error {
	var res V3Metro
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}