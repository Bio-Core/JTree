package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Experiment experiment
// swagger:model Experiment
type Experiment struct {

	// chip cartridge barcode
	ChipCartridgeBarcode *string `json:"experiments.chip_cartridge_barcode,omitempty" db:"experiments.chip_cartridge_barcode"`

	// complete date
	CompleteDate *string `json:"experiments.complete_date,omitempty" db:"experiments.complete_date"`

	// experiment id
	ExperimentID *string `json:"experiments.experiment_id,omitempty" db:"experiments.experiment_id"`

	// has project files
	HasProjectFiles *bool `json:"experiments.has_project_files,omitempty" db:"experiments.has_project_files"`

	// opened date
	OpenedDate *string `json:"experiments.opened_date,omitempty" db:"experiments.opened_date"`

	// panel assay screened
	PanelAssayScreened *int64 `json:"experiments.panel_assay_screened,omitempty" db:"experiments.panel_assay_screened"`

	// pcr
	Pcr *string `json:"experiments.pcr,omitempty" db:"experiments.pcr"`

	// priority
	Priority *int64 `json:"experiments.priority,omitempty" db:"experiments.priority"`

	// procedure order datetime
	ProcedureOrderDatetime *string `json:"experiments.procedure_order_datetime,omitempty" db:"experiments.procedure_order_datetime"`

	// project id
	ProjectID *string `json:"experiments.project_id,omitempty" db:"experiments.project_id"`

	// project name
	ProjectName *string `json:"experiments.project_name,omitempty" db:"experiments.project_name"`

	// sample id
	SampleID *string `json:"experiments.sample_id,omitempty" db:"experiments.sample_id"`

	// study id
	StudyID *string `json:"experiments.study_id,omitempty" db:"experiments.study_id"`

	// test date
	TestDate *string `json:"experiments.test_date,omitempty" db:"experiments.test_date"`
}

// Validate validates this experiment
func (m *Experiment) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Experiment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Experiment) UnmarshalBinary(b []byte) error {
	var res Experiment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
