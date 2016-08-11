package xacml

import "encoding/xml"

// 5.3: The <PolicyIssuer> element contains attributes describing the issuer of
// the policy or policy set. The use of the policy issuer element is defined in
// a separate administration profile [XACMLAdmin]. A PDP which does not
// implement the administration profile MUST report an error or return an
// Indeterminate result if it encounters this element.
type PolicyIssuer struct {
	XMLName xml.Name `xml:"PolicyIssuer"`

	// Free form XML describing the issuer. See Section 5.45.
	Content *Content `xml:"Content"`

	// An attribute of the issuer. See Section 5.46.
	Attributes []Attribute `xml:"Attribute"`
}

func (p PolicyIssuer) Validate(errs *Errors) {
	if p.Content != nil {
		p.Content.Validate(errs.Sub("Content"))
	}

	for i, el := range p.Attributes {
		el.Validate(errs.SubN("Attributes", i))
	}
}
