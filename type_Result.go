package xacml

import "encoding/xml"

// 5.48: The <Result> element represents an authorization decision result.  It
// MAY include a set of obligations that MUST be fulfilled by the PEP.  If the
// PEP does not understand or cannot fulfill an obligation, then the action of
// the PEP is determined by its bias, see section 7.1. It MAY include a set of
// advice with supplemental information which MAY be safely ignored by the PEP.
type Result struct {
	XMLName xml.Name `xml:"Result"`

	// The authorization decision: “Permit”, “Deny”, “Indeterminate” or
	// “NotApplicable”.
	Decision Decision `xml:"Decision"`

	// Indicates whether errors occurred during evaluation of the decision
	// request, and optionally, information about those errors.  If the
	// <Response> element contains <Result> elements whose <Status>
	// elements are all identical, and the <Response> element is contained
	// in a protocol wrapper that can convey status information, then the
	// common status information MAY be placed in the protocol wrapper and
	// this <Status> element MAY be omitted from all <Result> elements.
	Status *Status `xml:"Status"`

	// A list of obligations that MUST be fulfilled by the PEP.  If the PEP
	// does not understand or cannot fulfill an obligation, then the action
	// of the PEP is determined by its bias, see section 7.2.  See Section
	// 7.18 for a description of how the set of obligations to be returned
	// by the PDP is determined.
	Obligations *Obligations `xml:"Obligations"`

	// A list of advice that provide supplemental information to the PEP.
	// If the PEP does not understand an advice, the PEP may safely ignore
	// the advice. See Section 7.18 for a description of how the set of
	// advice to be returned by the PDP is determined.
	AssociatedAdvice *AssociatedAdvice `xml:"AssociatedAdvice"`

	// A list of attributes that were part of the request. The choice of
	// which attributes are included here is made with the IncludeInResult
	// attribute of the <Attribute> elements of the request. See section
	// 5.46.
	Attributes []Attributes `xml:"Attributes"`

	// If the ReturnPolicyIdList attribute in the <Request> is true (see
	// section 5.42), a PDP that implements this optional feature MUST
	// return a list of all policies which were found to be fully
	// applicable. That is, all policies where both the <Target> matched
	// and the <Condition> evaluated to true, whether or not the <Effect>
	// was the same or different from the <Decision>.
	PolicyIdentifierList *PolicyIdentifierList `xml:"PolicyIdentifierList"`
}

func (r Result) Validate(errs *Errors) {
	r.Decision.Validate(errs.Sub("Decision"))

	if r.Status != nil {
		r.Status.Validate(errs.Sub("Status"))
	}
	if r.Obligations != nil {
		r.Obligations.Validate(errs.Sub("Obligations"))
	}
	if r.AssociatedAdvice != nil {
		r.AssociatedAdvice.Validate(errs.Sub("AssociatedAdvice"))
	}
	if r.PolicyIdentifierList != nil {
		r.PolicyIdentifierList.Validate(errs.Sub("PolicyIdentifierList"))
	}

	for i, el := range r.Attributes {
		el.Validate(errs.SubN("Attributes", i))
	}
}
