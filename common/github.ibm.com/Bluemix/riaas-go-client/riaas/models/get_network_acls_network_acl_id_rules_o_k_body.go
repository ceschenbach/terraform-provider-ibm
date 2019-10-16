// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkAclsNetworkACLIDRulesOKBody TupleRuleCollection
// swagger:model getNetworkAclsNetworkAclIdRulesOKBody
type GetNetworkAclsNetworkACLIDRulesOKBody struct {

	// first
	First *GetNetworkAclsNetworkACLIDRulesOKBodyFirst `json:"first,omitempty"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// next
	Next *GetNetworkAclsNetworkACLIDRulesOKBodyNext `json:"next,omitempty"`

	// Collection of rules
	// Required: true
	Rules []*NetworkACLRule `json:"rules"`

	// The total number of resources across all pages
	// Minimum: 0
	TotalCount *int64 `json:"total_count,omitempty"`
}

// Validate validates this get network acls network Acl Id rules o k body
func (m *GetNetworkAclsNetworkACLIDRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBody) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBody) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", int64(m.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBody) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBody) validateTotalCount(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_count", "body", int64(*m.TotalCount), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNetworkAclsNetworkACLIDRulesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNetworkAclsNetworkACLIDRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkAclsNetworkACLIDRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
