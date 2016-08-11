package xacml

import "encoding/xml"

// 5.16: The <CombinerParameters> element conveys parameters for a policy- or
// rule-combining algorithm.
//
// If multiple <CombinerParameters> elements occur within the same policy or
// policy set, they SHALL be considered equal to one <CombinerParameters>
// element containing the concatenation of all the sequences of
// <CombinerParameters> contained in all the aforementioned
// <CombinerParameters> elements, such that the order of occurence of the
// <CominberParameters> elements is preserved in the concatenation of the
// <CombinerParameter> elements.
//
// Note that none of the combining algorithms specified in XACML 3.0 is
// parameterized.
type CombinerParameters struct {
	XMLName xml.Name `xml:"CombinerParameters"`

	// A single parameter.  See Section 5.17.
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}

func (c CombinerParameters) Validate(errs *Errors) {
	for i, el := range c.CombinerParameters {
		el.Validate(errs.SubN("CombinerParameters", i))
	}
}
