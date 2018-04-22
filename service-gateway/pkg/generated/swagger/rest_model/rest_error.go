// Code generated by go-swagger; DO NOT EDIT.

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestError rest error
// swagger:model RestError
type RestError struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this rest error
func (m *RestError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var restErrorTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InvalidSession","InternalServer","RecordDoesNotxist"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restErrorTypeTypePropEnum = append(restErrorTypeTypePropEnum, v)
	}
}

const (

	// RestErrorTypeInvalidSession captures enum value "InvalidSession"
	RestErrorTypeInvalidSession string = "InvalidSession"

	// RestErrorTypeInternalServer captures enum value "InternalServer"
	RestErrorTypeInternalServer string = "InternalServer"

	// RestErrorTypeRecordDoesNotxist captures enum value "RecordDoesNotxist"
	RestErrorTypeRecordDoesNotxist string = "RecordDoesNotxist"
)

// prop value enum
func (m *RestError) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, restErrorTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RestError) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestError) UnmarshalBinary(b []byte) error {
	var res RestError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}