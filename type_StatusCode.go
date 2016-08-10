package xacml

import "encoding/xml"

var kValidStatusCodes = []string{
	"urn:oasis:names:tc:xacml:1.0:status:ok",
	"urn:oasis:names:tc:xacml:1.0:status:missing-attribute",
	"urn:oasis:names:tc:xacml:1.0:status:syntax-error",
	"urn:oasis:names:tc:xacml:1.0:status:processing-error",
}

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
	} else {
		good := false
		for _, val := range kValidStatusCodes {
			if val == c.Value {
				good = true
				break
			}
		}

		if !good {
			errs.Addf("Invalid value for StatusCode: %s", c.Value)
		}
	}

	if c.StatusCode != nil {
		c.StatusCode.Validate(errs.Sub("StatusCode"))
	}
}
