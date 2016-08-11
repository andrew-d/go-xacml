package xacml

import "encoding/xml"

// 5.37: The <ObligationExpressions> element SHALL contain a set of
// <ObligationExpression> elements.
type ObligationExpressions struct {
	XMLName xml.Name `xml:"ObligationExpressions"`

	// A sequence of obligations expressions.  See Section 5.39.
	ObligationExpressions []ObligationExpression `xml:"ObligationExpression"`
}

func (o ObligationExpressions) Validate(errs *Errors) {
	if len(o.ObligationExpressions) < 1 {
		errs.Addf("ObligationExpressions not given")
	} else {
		for i, el := range o.ObligationExpressions {
			el.Validate(errs.SubN("ObligationExpressions", i))
		}
	}
}
