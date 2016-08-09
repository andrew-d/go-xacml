package xacml

import "encoding/xml"

// 5.55: The <StatusCode> element contains a major status code value and an
// optional sequence of minor status codes.
type StatusCode struct {
	XMLName xml.Name `xml:"StatusCode"`

	// See Section B.8 for a list of values.
	Value string `xml:",attr"`

	// Minor status code.  This status code qualifies its parent status
	// code.
	StatusCode *StatusCode
}

func (c StatusCode) Validate(errs *Errors) {
	if c.Value == "" {
		errs.Addf("Value not given")
	}

	if c.StatusCode != nil {
		c.StatusCode.Validate(errs.Sub("StatusCode"))
	}
}
