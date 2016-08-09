package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type PolicySetCombinerParameters struct {
	XMLName xml.Name `xml:"PolicySetCombinerParameters"`

	// Insert fields here
}

func (p PolicySetCombinerParameters) Validate(errs *Errors) {
}
