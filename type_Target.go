package xacml

import "encoding/xml"

// 5.6: The <Target> element identifies the set of decision requests that the
// parent element is intended to evaluate.  The <Target> element SHALL appear
// as a child of a <PolicySet> and <Policy> element and MAY appear as a child
// of a <Rule> element.
//
// The <Target> element SHALL contain a conjunctive sequence of <AnyOf>
// elements.  For the parent of the <Target> element to be applicable to the
// decision request, there MUST be at least one positive match between each
// <AnyOf> element of the <Target> element and the corresponding section of the
// <Request> element.
type Target struct {
	XMLName xml.Name `xml:"Target"`

	// Matching specification for attributes in the context.  If this element
	// is missing, then the target SHALL match all contexts.
	AnyOf []AnyOf `xml:"AnyOf"`
}

func (t Target) Validate(errs *Errors) {
	for i, el := range t.AnyOf {
		el.Validate(errs.SubN("AnyOf", i))
	}
}
