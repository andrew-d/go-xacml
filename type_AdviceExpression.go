package xacml

import "encoding/xml"

// 5.40: The <AdviceExpression> element evaluates to an advice and SHALL
// contain an identifier for an advice and a set of expressions that form
// arguments of the supplemental information defined by the advice.  The
// AppliesTo attribute SHALL indicate the effect for which this advice must be
// provided to the PEP.
type AdviceExpression struct {
	XMLName xml.Name `xml:"AdviceExpression"`

	// Advice identifier.  The value of the advice identifier MAY be
	// interpreted by the PEP.
	AdviceId string `xml:",attr"`

	// The effect for which this advice must be provided to the PEP.
	AppliesTo string `xml:",attr"`

	// Advice arguments in the form of expressions. The expressions SHALL
	// be evaluated by the PDP to constant <AttributeValue> elements or
	// bags, which shall be the attribute assignments in the <Advice>
	// returned to the PEP.  If an <AttributeAssignmentExpression>
	// evaluates to an atomic attribute value, then there MUST be one
	// resulting <AttributeAssignment> which MUST contain this single
	// attribute value. If the <AttributeAssignmentExpression> evaluates to
	// a bag, then there MUST be a resulting <AttributeAssignment> for each
	// of the values in the bag. If the bag is empty, there shall be no
	// <AttributeAssignment> from this <AttributeAssignmentExpression>.
	// The values of the advice arguments MAY be interpreted by the PEP.
	AttributeAssignmentExpressions []AttributeAssignmentExpression `xml:"AttributeAssignmentExpression"`
}

func (a AdviceExpression) Validate(errs *Errors) {
	if a.AdviceId == "" {
		errs.Addf("AdviceId not given")
	}
	if a.AppliesTo == "" {
		errs.Addf("AppliesTo not given")
	} else if a.AppliesTo != "Permit" && a.AppliesTo != "Deny" {
		errs.Addf("Invalid value for AppliesTo: %q", a.AppliesTo)
	}

	for i, el := range a.AttributeAssignmentExpressions {
		el.Validate(errs.SubN("AttributeAssignmentExpressions", i))
	}
}
