package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type CombinerParameters struct {
	XMLName xml.Name `xml:"CombinerParameters"`

	// Insert fields here
}

func (c CombinerParameters) Validate(errs *Errors) {
}
