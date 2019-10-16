// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile reference
// swagger:model postInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile
type PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile struct {

	// The CRN for this profile
	Crn string `json:"crn,omitempty"`

	// The user-defined name for this resource
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post instances params body volume attachments items volume profile
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile) UnmarshalBinary(b []byte) error {
	var res PostInstancesParamsBodyVolumeAttachmentsItemsVolumeProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
