package xacml

import "encoding/xml"

// 5.34: The <Obligation> element SHALL contain an identifier for the
// obligation and a set of attributes that form arguments of the action defined
// by the obligation.
type Obligation struct {
	XMLName xml.Name `xml:"Obligation"`

	// Obligation identifier.  The value of the obligation identifier SHALL
	// be interpreted by the PEP.
	ObligationId string `xml:",attr"`

	// The spec does not include this attribute, but a bunch of our
	// conformance tests do.  We're going to include it anyway.
	FulfillOn string `xml:",attr,omitempty"`

	// Obligation arguments assignment.  The values of the obligation
	// arguments SHALL be interpreted by the PEP.
	AttributeAssignments []AttributeAssignment `xml:"AttributeAssignment"`
}

func (o Obligation) Validate(errs *Errors) {
	if o.ObligationId == "" {
		errs.Addf("ObligationId not given")
	}

	for i, el := range o.AttributeAssignments {
		el.Validate(errs.SubN("AttributeAssignments", i))
	}
}
