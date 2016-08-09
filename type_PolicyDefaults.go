package xacml

import "encoding/xml"

// 5.15: The <PolicyDefaults> element SHALL specify default values that apply
// to the <Policy> element.
type PolicyDefaults struct {
	XMLName xml.Name `xml:"PolicyDefaults"`

	// Default XPath version.
	XPathVersion XPathVersion `xml:"XPathVersion"`
}

func (p PolicyDefaults) Validate(errs *Errors) {
	p.XPathVersion.Validate(errs.Sub("XPathVersion"))
}
