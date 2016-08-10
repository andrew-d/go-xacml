package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type MultiRequests struct {
	XMLName xml.Name `xml:"MultiRequests"`

	// Insert fields here
}

func (m MultiRequests) Validate(errs *Errors) {
}
