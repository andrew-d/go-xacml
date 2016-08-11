package xacml

import "encoding/xml"

// 5.38: The <AdviceExpressions> element SHALL contain a set of
// <AdviceExpression> elements.
type AdviceExpressions struct {
	XMLName xml.Name `xml:"AdviceExpressions"`

	// A sequence of advice expressions.  See Section 5.40.
	AdviceExpressions []AdviceExpression `xml:"AdviceExpression"`
}

func (a AdviceExpressions) Validate(errs *Errors) {
	if len(a.AdviceExpressions) < 1 {
		errs.Addf("AdviceExpressions not given")
	} else {
		for i, el := range a.AdviceExpressions {
			el.Validate(errs.SubN("AdviceExpressions", i))
		}
	}
}
