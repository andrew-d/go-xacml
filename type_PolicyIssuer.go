package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type PolicyIssuer struct {
	XMLName xml.Name `xml:"PolicyIssuer"`

	// Insert fields here
}

func (p PolicyIssuer) Validate(errs *Errors) {
}
