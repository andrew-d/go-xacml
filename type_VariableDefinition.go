package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type VariableDefinition struct {
	XMLName xml.Name `xml:"VariableDefinition"`

	// Insert fields here
}

func (v VariableDefinition) Validate(errs *Errors) {
}
