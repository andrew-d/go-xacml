package xacml

import "encoding/xml"

// 5.18: The <RuleCombinerParameters> element conveys parameters associated
// with a particular rule within a policy for a rule-combining algorithm.
//
// Each <RuleCombinerParameters> element MUST be associated with a rule
// contained within the same policy.  If multiple <RuleCombinerParameters>
// elements reference the same rule, they SHALL be considered equal to one
// <RuleCombinerParameters> element containing the concatenation of all the
// sequences of <CombinerParameters> contained in all the aforementioned
// <RuleCombinerParameters> elements, such that the order of occurrence of the
// <RuleCominberParameters> elements is preserved in the concatenation of the
// <CombinerParameter> elements.
//
// Note that none of the rule-combining algorithms specified in XACML 3.0 is
// parameterized.
//
// Support for the <RuleCombinerParameters> element is optional, only if
// support for combiner parameters is not implemented.
type RuleCombinerParameters struct {
	XMLName xml.Name `xml:"RuleCombinerParameters"`

	// The identifier of the <Rule> contained in the policy.
	RuleIdRef string `xml:",attr"`

	// A single parameter.  See Section 5.17.
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}

func (r RuleCombinerParameters) Validate(errs *Errors) {
	if r.RuleIdRef == "" {
		errs.Addf("RuleIdRef not given")
	}

	for i, el := range r.CombinerParameters {
		el.Validate(errs.SubN("CombinerParameters", i))
	}
}
