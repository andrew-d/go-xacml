package xacml

import "encoding/xml"

// 5.54: The <Status> element represents the status of the authorization
// decision result.
type Status struct {
	XMLName xml.Name `xml:"Status"`

	// Status code.
	StatusCode StatusCode `xml:"StatusCode"`

	// A status message describing the status code.
	StatusMessage string `xml:"StatusMessage,omitempty"`

	// Additional status information.
	StatusDetail *StatusDetail `xml:"StatusDetail"`
}

func (s Status) Validate(errs *Errors) {
	s.StatusCode.Validate(errs.Sub("StatusCode"))

	if s.StatusDetail != nil {
		s.StatusDetail.Validate(errs.Sub("StatusDetail"))
	}
}
