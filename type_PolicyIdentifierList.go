package xacml

import "encoding/xml"

// 5.49: The <PolicyIdentifierList> element contains a list of policy and
// policy set identifiers of policies which have been applicable to a request.
// The list is unordered.
type PolicyIdentifierList struct {
	XMLName xml.Name `xml:"PolicyIdentifierList"`

	// The identifier and version of a policy which was applicable to the
	// request. See section 5.11. The <PolicyIdReference> element MUST use
	// the Version attribute to specify the version and MUST NOT use the
	// LatestVersion or EarliestVersion attributes.
	PolicyIdReferences []PolicyIdReference `xml:"PolicyIdReference"`

	// The identifier and version of a policy set which was applicable to
	// the request. See section 5.10. The <PolicySetIdReference> element
	// MUST use the Version attribute to specify the version and MUST NOT
	// use the LatestVersion or EarliestVersion attributes.
	PolicySetIdReferences []PolicySetIdReference `xml:"PolicySetIdReference"`
}

func (p PolicyIdentifierList) Validate(errs *Errors) {
	for i, el := range p.PolicyIdReferences {
		el.Validate(errs.SubN("PolicyIdReferences", i))
	}
	for i, el := range p.PolicySetIdReferences {
		el.Validate(errs.SubN("PolicySetIdReferences", i))
	}
}
