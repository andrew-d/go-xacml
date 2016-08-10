package xacml

import "encoding/xml"

// 5.57: The <StatusDetail> element qualifies the <Status> element with
// additional information.
type StatusDetail struct {
	XMLName xml.Name `xml:"StatusDetail"`

	// The <StatusDetail> element allows arbitrary XML content.
	XML string `xml:",innerxml"`
}

func (s StatusDetail) Validate(errs *Errors) {
	// Validations require knowing the status code, which we don't.
}
