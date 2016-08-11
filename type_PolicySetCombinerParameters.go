package xacml

import "encoding/xml"

// 5.20: The <PolicySetCombinerParameters> element conveys parameters
// associated with a particular policy set within a policy set for a
// policy-combining algorithm.
//
// Each <PolicySetCombinerParameters> element MUST be associated with a policy
// set contained within the same policy set.  If multiple
// <PolicySetCombinerParameters> elements reference the same policy set, they
// SHALL be considered equal to one <PolicySetCombinerParameters> element
// containing the concatenation of all the sequences of <CombinerParameters>
// contained in all the aforementioned <PolicySetCombinerParameters> elements,
// such that the order of occurrence of the <PolicySetCominberParameters>
// elements is preserved in the concatenation of the <CombinerParameter>
// elements.
//
// Note that none of the policy-combining algorithms specified in XACML 3.0 is
// parameterized.
//
// Support for the <PolicySetCombinerParameters> element is optional, only if
// support for combiner parameters is not implemented.
type PolicySetCombinerParameters struct {
	XMLName xml.Name `xml:"PolicySetCombinerParameters"`

	// The identifier of a <PolicySet> or the value of a
	// <PolicySetIdReference> contained in the policy set.
	PolicySetIdRef string `xml:",attr"`

	// A single parameter.  See Section 5.17.
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}

func (p PolicySetCombinerParameters) Validate(errs *Errors) {
	if p.PolicySetIdRef == "" {
		errs.Addf("PolicySetIdRef not given")
	}

	for i, el := range p.CombinerParameters {
		el.Validate(errs.SubN("CombinerParameters", i))
	}
}
