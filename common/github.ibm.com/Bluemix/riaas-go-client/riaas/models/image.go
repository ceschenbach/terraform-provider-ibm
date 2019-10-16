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

// Image Image
// swagger:model image
type Image struct {

	// The image architecture
	// Enum: [amd64 powerpc]
	Architecture string `json:"architecture,omitempty"`

	// The date and time that the image was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this image
	Crn string `json:"crn,omitempty"`

	// file
	File *ImageFile `json:"file,omitempty"`

	// format
	// Enum: [raw qcow2 vmdk vhdx vdi ova box]
	Format string `json:"format,omitempty"`

	// The URL for this image
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this image
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this image
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// operating system
	OperatingSystem *ImageOperatingSystem `json:"operating_system,omitempty"`

	// resource group
	ResourceGroup *ResourceReferenceWithOutCrn `json:"resource_group,omitempty"`

	// status
	// Enum: [pending available corrupt]
	Status string `json:"status,omitempty"`

	// Whether the image is publicly visible or private to the account
	// Enum: [public private]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
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

	if err := m.validateOperatingSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var imageTypeArchitecturePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["amd64","powerpc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeArchitecturePropEnum = append(imageTypeArchitecturePropEnum, v)
	}
}

const (

	// ImageArchitectureAmd64 captures enum value "amd64"
	ImageArchitectureAmd64 string = "amd64"

	// ImageArchitecturePowerpc captures enum value "powerpc"
	ImageArchitecturePowerpc string = "powerpc"
)

// prop value enum
func (m *Image) validateArchitectureEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, imageTypeArchitecturePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Image) validateArchitecture(formats strfmt.Registry) error {

	if swag.IsZero(m.Architecture) { // not required
		return nil
	}

	// value enum
	if err := m.validateArchitectureEnum("architecture", "body", m.Architecture); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

var imageTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["raw","qcow2","vmdk","vhdx","vdi","ova","box"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeFormatPropEnum = append(imageTypeFormatPropEnum, v)
	}
}

const (

	// ImageFormatRaw captures enum value "raw"
	ImageFormatRaw string = "raw"

	// ImageFormatQcow2 captures enum value "qcow2"
	ImageFormatQcow2 string = "qcow2"

	// ImageFormatVMDK captures enum value "vmdk"
	ImageFormatVMDK string = "vmdk"

	// ImageFormatVhdx captures enum value "vhdx"
	ImageFormatVhdx string = "vhdx"

	// ImageFormatVdi captures enum value "vdi"
	ImageFormatVdi string = "vdi"

	// ImageFormatOva captures enum value "ova"
	ImageFormatOva string = "ova"

	// ImageFormatBox captures enum value "box"
	ImageFormatBox string = "box"
)

// prop value enum
func (m *Image) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, imageTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Image) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateOperatingSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatingSystem) { // not required
		return nil
	}

	if m.OperatingSystem != nil {
		if err := m.OperatingSystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operating_system")
			}
			return err
		}
	}

	return nil
}

func (m *Image) validateResourceGroup(formats strfmt.Registry) error {

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

var imageTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","available","corrupt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeStatusPropEnum = append(imageTypeStatusPropEnum, v)
	}
}

const (

	// ImageStatusPending captures enum value "pending"
	ImageStatusPending string = "pending"

	// ImageStatusAvailable captures enum value "available"
	ImageStatusAvailable string = "available"

	// ImageStatusCorrupt captures enum value "corrupt"
	ImageStatusCorrupt string = "corrupt"
)

// prop value enum
func (m *Image) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, imageTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Image) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var imageTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeVisibilityPropEnum = append(imageTypeVisibilityPropEnum, v)
	}
}

const (

	// ImageVisibilityPublic captures enum value "public"
	ImageVisibilityPublic string = "public"

	// ImageVisibilityPrivate captures enum value "private"
	ImageVisibilityPrivate string = "private"
)

// prop value enum
func (m *Image) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, imageTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Image) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
