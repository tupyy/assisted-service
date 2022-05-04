// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ImageType image type
//
// swagger:model image_type
type ImageType string

func NewImageType(value ImageType) *ImageType {
	v := value
	return &v
}

const (

	// ImageTypeFullIso captures enum value "full-iso"
	ImageTypeFullIso ImageType = "full-iso"

	// ImageTypeMinimalIso captures enum value "minimal-iso"
	ImageTypeMinimalIso ImageType = "minimal-iso"
)

// for schema
var imageTypeEnum []interface{}

func init() {
	var res []ImageType
	if err := json.Unmarshal([]byte(`["full-iso","minimal-iso"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeEnum = append(imageTypeEnum, v)
	}
}

func (m ImageType) validateImageTypeEnum(path, location string, value ImageType) error {
	if err := validate.EnumCase(path, location, value, imageTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this image type
func (m ImageType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImageTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this image type based on context it is used
func (m ImageType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
