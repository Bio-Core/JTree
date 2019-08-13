package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Patient patient
// swagger:model Patient
type Patient struct {

	// clinical history
	ClinicalHistory *string `json:"patients.clinical_history,omitempty" db:"patients.clinical_history"`

	// date received
	DateReceived *time.Time `json:"patients.date_received,omitempty" db:"patients.date_received"`

	// date reported
	DateReported *time.Time `json:"patients.date_reported,omitempty" db:"patients.date_reported"`

	// dob
	Dob *time.Time `json:"patients.dob,omitempty" db:"patients.dob"`

	// first name
	FirstName *string `json:"patients.first_name,omitempty" db:"patients.first_name"`

	// gender
	Gender *string `json:"patients.gender,omitempty" db:"patients.gender"`

	// initials
	Initials *string `json:"patients.initials,omitempty" db:"patients.initials"`

	// last name
	LastName *string `json:"patients.last_name,omitempty" db:"patients.last_name"`

	// mrn
	Mrn *string `json:"patients.mrn,omitempty" db:"patients.mrn"`

	// on hcn
	OnHcn *string `json:"patients.on_hcn,omitempty" db:"patients.on_hcn"`

	// patient id
	PatientID *string `json:"patients.patient_id,omitempty" db:"patients.patient_id"`

	// patient type
	PatientType *string `json:"patients.patient_type,omitempty" db:"patients.patient_type"`

	// referring physician
	ReferringPhysican *string `json:"patients.referring_physician,omitempty" db:"patients.referring_physician"`

	// se num
	SeNum *string `json:"patients.se_num,omitempty" db:"patients.se_num"`

	// surgical date
	SurgicalDate *time.Time `json:"patients.surgical_date,omitempty" db:"patients.surgical_date"`
}

// Validate validates this patient
func (m *Patient) Validate(formats strfmt.Registry) error {
	var res []error
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Patient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Patient) UnmarshalBinary(b []byte) error {
	var res Patient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
