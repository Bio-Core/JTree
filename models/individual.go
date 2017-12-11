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

// Individual individual
// swagger:model Individual
type Individual struct {

	// attributes
	// Required: true
	Attributes Attributes `json:"attributes"`

	// created date
	// Required: true
	CreatedDate *strfmt.Date `json:"createdDate"`

	// description
	// Required: true
	Description *string `json:"description"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sex
	// Required: true
	Sex *OntologyTerm `json:"sex"`

	// species
	// Required: true
	Species *OntologyTerm `json:"species"`

	// updated date
	// Required: true
	UpdatedDate *strfmt.Date `json:"updatedDate"`
}

// Validate validates this individual
func (m *Individual) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSex(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpecies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Individual) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if err := m.Attributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes")
		}
		return err
	}

	return nil
}

func (m *Individual) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("createdDate", "body", m.CreatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("createdDate", "body", "date", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Individual) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Individual) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Individual) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Individual) validateSex(formats strfmt.Registry) error {

	if err := validate.Required("sex", "body", m.Sex); err != nil {
		return err
	}

	if m.Sex != nil {

		if err := m.Sex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sex")
			}
			return err
		}
	}

	return nil
}

func (m *Individual) validateSpecies(formats strfmt.Registry) error {

	if err := validate.Required("species", "body", m.Species); err != nil {
		return err
	}

	if m.Species != nil {

		if err := m.Species.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("species")
			}
			return err
		}
	}

	return nil
}

func (m *Individual) validateUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("updatedDate", "body", m.UpdatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedDate", "body", "date", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Individual) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Individual) UnmarshalBinary(b []byte) error {
	var res Individual
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
