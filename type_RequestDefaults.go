package xacml

import "encoding/xml"

// 5.43: The <RequestDefaults> element SHALL specify default values that apply
// to the <Request> element.
type RequestDefaults struct {
	XMLName xml.Name `xml:"RequestDefaults"`

	// Default XPath version for XPath expressions occurring in the request.
	XPathVersion XPathVersion `xml:"XPathVersion"`
}

func (r RequestDefaults) Validate(errs *Errors) {
	r.XPathVersion.Validate(errs.Sub("XPathVersion"))
}
