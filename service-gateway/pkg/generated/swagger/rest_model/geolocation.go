// Code generated by go-swagger; DO NOT EDIT.

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Geolocation geolocation
// swagger:model Geolocation
type Geolocation struct {

	// api desired accuracy
	APIDesiredAccuracy int32 `json:"apiDesiredAccuracy,omitempty"`

	// api language
	APILanguage string `json:"apiLanguage,omitempty"`

	// api provider
	APIProvider string `json:"apiProvider,omitempty"`

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// common name
	CommonName string `json:"commonName,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// country code
	CountryCode string `json:"countryCode,omitempty"`

	// formatted address
	FormattedAddress string `json:"formattedAddress,omitempty"`

	// google place ID
	GooglePlaceID string `json:"googlePlaceID,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// latitude
	Latitude float32 `json:"latitude,omitempty"`

	// longitude
	Longitude float32 `json:"longitude,omitempty"`

	// neighborhood
	Neighborhood string `json:"neighborhood,omitempty"`

	// postal code
	PostalCode string `json:"postalCode,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// route
	Route string `json:"route,omitempty"`

	// session ID
	SessionID string `json:"sessionID,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// state code
	StateCode string `json:"stateCode,omitempty"`

	// street
	Street string `json:"street,omitempty"`

	// street number
	StreetNumber string `json:"streetNumber,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// town
	Town string `json:"town,omitempty"`
}

// Validate validates this geolocation
func (m *Geolocation) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Geolocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Geolocation) UnmarshalBinary(b []byte) error {
	var res Geolocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}