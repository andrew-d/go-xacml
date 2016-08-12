package xacml

import "encoding/xml"

// 5.21: The <Rule> element SHALL define the individual rules in the policy.
// The main components of this element are the <Target>, <Condition>,
// <ObligationExpressions>  and <AdviceExpressions>  elements and the Effect
// attribute.
//
// A <Rule> element may be evaluated, in which case the evaluation procedure
// defined in Section 7.10 SHALL be used.
type Rule struct {
	XMLName xml.Name `xml:"Rule"`

	// A string identifying this rule.
	RuleId string `xml:",attr"`

	// Rule effect.  The value of this attribute is either “Permit” or “Deny”.
	Effect string `xml:",attr"`

	// A free-form description of the rule.
	Description string `xml:"Description,omitempty"`

	// Identifies the set of decision requests that the <Rule> element is
	// intended to evaluate.  If this element is omitted, then the target for
	// the <Rule> SHALL be defined by the <Target> element of the enclosing
	// <Policy> element.  See Section 7.7 for details.
	Target *Target `xml:"Target"`

	// A predicate that MUST be satisfied for the rule to be assigned its
	// Effect value.
	Condition *Condition `xml:"Condition"`

	// A conjunctive sequence of obligation expressions which MUST be evaluated
	// into obligations byt the PDP. The corresponsding obligations MUST be
	// fulfilled by the PEP in conjunction with the authorization decision.
	// See Section 7.18 for a description of how the set of obligations to be
	// returned by the PDP SHALL be determined. See section 7.2 about
	// enforcement of obligations.
	ObligationExpressions *ObligationExpressions `xml:"ObligationExpressions"`

	// A conjunctive sequence of advice expressions which MUST evaluated into
	// advice by the PDP. The corresponding advice provide supplementary
	// information to the PEP in conjunction with the authorization decision.
	// See Section 7.18 for a description of how the set of advice to be
	// returned by the PDP SHALL be determined.
	AdviceExpressions *AdviceExpressions `xml:"AdviceExpressions"`
}

func (r Rule) Validate(errs *Errors) {
	if r.RuleId == "" {
		errs.Addf("RuleId not given")
	}
	if r.Effect != "Permit" && r.Effect != "Deny" {
		errs.Addf("Effect must be \"Permit\" or \"Deny\", not %q", r.Effect)
	}

	// Optional
	if r.Target != nil {
		r.Target.Validate(errs.Sub("Target"))
	}

	if r.Condition != nil {
		r.Condition.Validate(errs.Sub("Condition"))
	}

	if r.ObligationExpressions != nil {
		r.ObligationExpressions.Validate(errs.Sub("ObligationExpressions"))
	}

	if r.AdviceExpressions != nil {
		r.AdviceExpressions.Validate(errs.Sub("AdviceExpressions"))
	}
}
