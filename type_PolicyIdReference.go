package xacml

import "encoding/xml"

// 5.11: The <PolicyIdReference> element SHALL be used to reference a <Policy>
// element by id.  If <PolicyIdReference> is a URL, then it MAY be resolvable
// to the <Policy> element.  However, the mechanism for resolving a policy
// reference to the corresponding policy is outside the scope of this
// specification.
type PolicyIdReference struct {
	XMLName xml.Name `xml:"PolicyIdReference"`

	// Specifies a matching expression for the version of the policy
	// referenced.
	Version string `xml:",attr"`

	// Specifies a matching expression for the earliest acceptable version
	// of the policy referenced.
	EarliestVersion string `xml:",attr"`

	// Specifies a matching expression for the latest acceptable version of
	// the policy referenced.
	LatestVersion string `xml:",attr"`

	// TODO: Do we need to use "innerxml", or can we actually use "chardata"?
	Value string `xml:",chardata"`
}

func (p PolicyIdReference) Validate(errs *Errors) {
	if p.Value == "" {
		errs.Addf("Value not given")
	}
}
