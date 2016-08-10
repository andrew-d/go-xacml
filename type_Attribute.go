package xacml

import "encoding/xml"

// 5.46: The <Attribute> element is the central abstraction of the request
// context.  It contains attribute meta-data and one or more attribute values.
// The attribute meta-data comprises the attribute identifier and the attribute
// issuer.  <AttributeDesignator> elements in the policy MAY refer to
// attributes by means of this meta-data.
type Attribute struct {
	XMLName xml.Name `xml:"Attribute"`

	// The Attribute identifier.  A number of identifiers are reserved by
	// XACML to denote commonly used attributes.  See Appendix B.
	AttributeId string `xml:",attr"`

	// The Attribute issuer.  For example, this attribute value MAY be an
	// x500Name that binds to a public key, or it may be some other
	// identifier exchanged out-of-band by issuing and relying parties.
	Issuer *string `xml:",attr"`

	// Whether to include this attribute in the result. This is useful to
	// correlate requests with their responses in case of multiple
	// requests.
	// NOTE: Optional, defaults to false
	IncludeInResult bool `xml:",attr"`

	// One or more attribute values.  Each attribute value MAY have
	// contents that are empty, occur once or occur multiple times.
	AttributeValues []AttributeValue `xml:"AttributeValue"`
}

func (a Attribute) Validate(errs *Errors) {
	if a.AttributeId == "" {
		errs.Addf("AttributeId not given")
	}

	if a.Issuer != nil && *a.Issuer == "" {
		errs.Addf("Issuer cannot be empty")
	}

	if len(a.AttributeValues) < 1 {
		errs.Addf("AttributeValues not given")
	} else {
		for i, el := range a.AttributeValues {
			el.Validate(errs.SubN("AttributeValues", i))
		}
	}
}
