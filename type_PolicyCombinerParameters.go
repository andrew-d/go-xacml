package xacml

import "encoding/xml"

// 5.19: The <PolicyCombinerParameters> element conveys parameters associated
// with a particular policy within a policy set for a policy-combining
// algorithm.
//
// Each <PolicyCombinerParameters> element MUST be associated with a policy
// contained within the same policy set.  If multiple
// <PolicyCombinerParameters> elements reference the same policy, they SHALL be
// considered equal to one <PolicyCombinerParameters> element containing the
// concatenation of all the sequences of <CombinerParameters> contained in all
// the aforementioned <PolicyCombinerParameters> elements, such that the order
// of occurrence of the <PolicyCominberParameters> elements is preserved in the
// concatenation of the <CombinerParameter> elements.
//
// Note that none of the policy-combining algorithms specified in XACML 3.0 is
// parameterized.
//
// Support for the <PolicyCombinerParameters> element is optional, only if
// support for combiner parameters is not implemented.
type PolicyCombinerParameters struct {
	XMLName xml.Name `xml:"PolicyCombinerParameters"`

	// The identifier of a <Policy> or the value of a <PolicyIdReference>
	// contained in the policy set.
	PolicyIdRef string `xml:",attr"`

	// A single parameter.  See Section 5.17.
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}

func (p PolicyCombinerParameters) Validate(errs *Errors) {
	if p.PolicyIdRef == "" {
		errs.Addf("PolicyIdRef not given")
	}

	for i, el := range p.CombinerParameters {
		el.Validate(errs.SubN("CombinerParameters", i))
	}
}
