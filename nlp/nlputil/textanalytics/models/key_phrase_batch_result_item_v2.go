package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// KeyPhraseBatchResultItemV2 key phrase batch result item v2
// swagger:model KeyPhraseBatchResultItemV2
type KeyPhraseBatchResultItemV2 struct {

	// Unique document identifier.
	ID string `json:"id,omitempty"`

	// A list of representative words or phrases. The number of key phrases returned is proportional to the number of words in the input document.
	KeyPhrases []string `json:"keyPhrases"`
}

// Validate validates this key phrase batch result item v2
func (m *KeyPhraseBatchResultItemV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyPhrases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyPhraseBatchResultItemV2) validateKeyPhrases(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyPhrases) { // not required
		return nil
	}

	return nil
}