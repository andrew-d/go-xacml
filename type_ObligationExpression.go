package xacml

import "encoding/xml"

// 5.39: The <ObligationExpression> element evaluates to an obligation and
// SHALL contain an identifier for an obligation and a set of expressions that
// form arguments of the action defined by the obligation.  The FulfillOn
// attribute SHALL indicate the effect for which this obligation must be
// fulfilled by the PEP.
type ObligationExpression struct {
	XMLName xml.Name `xml:"ObligationExpression"`

	// Obligation identifier.  The value of the obligation identifier SHALL
	// be interpreted by the PEP.
	ObligationId string `xml:",attr"`

	// The effect for which this obligation must be fulfilled by the PEP.
	FulfillOn string `xml:",attr"`

	// Obligation arguments in the form of expressions. The expressions
	// SHALL be evaluated by the PDP to constant <AttributeValue> elements
	// or bags, which shall be the attribute assignments in the
	// <Obligation> returned to the PEP. If an
	// <AttributeAssignmentExpression> evaluates to an atomic attribute
	// value, then there MUST be one resulting <AttributeAssignment> which
	// MUST contain this single attribute value. If the
	// <AttributeAssignmentExpression> evaluates to a bag, then there MUST
	// be a resulting <AttributeAssignment> for each of the values in the
	// bag. If the bag is empty, there shall be no <AttributeAssignment>
	// from this <AttributeAssignmentExpression>.The values of the
	// obligation arguments SHALL be interpreted by the PEP.
	AttributeAssignmentExpressions []AttributeAssignmentExpression `xml:"AttributeAssignmentExpression"`
}

func (o ObligationExpression) Validate(errs *Errors) {
	if o.ObligationId == "" {
		errs.Addf("ObligationId not given")
	}
	if o.FulfillOn == "" {
		errs.Addf("FulfillOn not given")
	} else if o.FulfillOn != "Permit" && o.FulfillOn != "Deny" {
		errs.Addf("Invalid value for FulfillOn: %q", o.FulfillOn)
	}

	for i, el := range o.AttributeAssignmentExpressions {
		el.Validate(errs.SubN("AttributeAssignmentExpressions", i))
	}
}
