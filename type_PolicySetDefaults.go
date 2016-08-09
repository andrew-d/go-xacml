package xacml

import "encoding/xml"

// 5.4: The <PolicySetDefaults> element SHALL specify default values that apply
// to the <PolicySet> element.
type PolicySetDefaults struct {
	XMLName xml.Name `xml:"PolicySetDefaults"`

	// Default XPath version.
	XPathVersion XPathVersion `xml:"XPathVersion"`
}

func (p PolicySetDefaults) Validate(errs *Errors) {
	p.XPathVersion.Validate(errs.Sub("XPathVersion"))
}
