// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Vpc VPC
// swagger:model vpc
type Vpc struct {

	// Indicates whether this VPC is connected to Classic Infrastructure. If true, this VPC's
	// resources have private network connectivity to the account's Classice Infrastructure
	// resources. Only one VPC on an account may be connected in this way. This value is set at
	// creation and subsequently immutable.
	//
	ClassicAccess bool `json:"classic_access,omitempty"`

	// The date and time that the VPC was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this VPC
	Crn string `json:"crn,omitempty"`

	// default network acl
	DefaultNetworkACL *ResourceReference `json:"default_network_acl,omitempty"`

	// default security group
	DefaultSecurityGroup *ResourceReference `json:"default_security_group,omitempty"`

	// The URL for this VPC
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this VPC
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this VPC
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// status
	// Enum: [available pending deleting failed]
	Status string `json:"status,omitempty"`
}

// Validate validates this vpc
func (m *Vpc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultNetworkACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vpc) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Vpc) validateDefaultNetworkACL(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultNetworkACL) { // not required
		return nil
	}

	if m.DefaultNetworkACL != nil {
		if err := m.DefaultNetworkACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_network_acl")
			}
			return err
		}
	}

	return nil
}

func (m *Vpc) validateDefaultSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultSecurityGroup) { // not required
		return nil
	}

	if m.DefaultSecurityGroup != nil {
		if err := m.DefaultSecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_security_group")
			}
			return err
		}
	}

	return nil
}

func (m *Vpc) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Vpc) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Vpc) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Vpc) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

var vpcTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","pending","deleting","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpcTypeStatusPropEnum = append(vpcTypeStatusPropEnum, v)
	}
}

const (

	// VpcStatusAvailable captures enum value "available"
	VpcStatusAvailable string = "available"

	// VpcStatusPending captures enum value "pending"
	VpcStatusPending string = "pending"

	// VpcStatusDeleting captures enum value "deleting"
	VpcStatusDeleting string = "deleting"

	// VpcStatusFailed captures enum value "failed"
	VpcStatusFailed string = "failed"
)

// prop value enum
func (m *Vpc) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpcTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Vpc) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vpc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vpc) UnmarshalBinary(b []byte) error {
	var res Vpc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
