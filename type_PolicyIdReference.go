package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type PolicyIdReference struct {
	XMLName xml.Name `xml:"PolicyIdReference"`

	// Insert fields here
}

func (p PolicyIdReference) Validate(errs *Errors) {
}
