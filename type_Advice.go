package xacml

import "encoding/xml"

// 5.35: The <Advice> element SHALL contain an identifier for the advice and a
// set of attributes that form arguments of the supplemental information
// defined by the advice.
type Advice struct {
	XMLName xml.Name `xml:"Advice"`

	// Advice identifier.  The value of the advice identifier MAY be
	// interpreted by the PEP.
	AdviceId string `xml:",attr"`

	// Advice arguments assignment.  The values of the advice arguments MAY be
	// interpreted by the PEP.
	Assignments []AttributeAssignment `xml:"AttributeAssignment"`
}

func (a Advice) Validate(errs *Errors) {
	if a.AdviceId == "" {
		errs.Addf("AdviceId not given")
	}

	for i, assignment := range a.Assignments {
		assignment.Validate(errs.SubN("Assignments", i))
	}
}
