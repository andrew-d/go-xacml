package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type PolicySetIdReference struct {
	XMLName xml.Name `xml:"PolicySetIdReference"`

	// Insert fields here
}

func (p PolicySetIdReference) Validate(errs *Errors) {
}
