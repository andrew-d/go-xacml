package xacml

import "encoding/xml"

// 5.10: The <PolicySetIdReference> element SHALL be used to reference a
// <PolicySet> element by id.  If <PolicySetIdReference> is a URL, then it MAY
// be resolvable to the <PolicySet> element.  However, the mechanism for
// resolving a policy set reference to the corresponding policy set is outside
// the scope of this specification.
type PolicySetIdReference struct {
	XMLName xml.Name `xml:"PolicySetIdReference"`

	// Specifies a matching expression for the version of the policy set
	// referenced.
	Version string `xml:",attr,omitempty"`

	// Specifies a matching expression for the earliest acceptable version
	// of the policy set referenced.
	EarliestVersion string `xml:",attr,omitempty"`

	// Specifies a matching expression for the latest acceptable version of
	// the policy set referenced.
	LatestVersion string `xml:",attr,omitempty"`

	// TODO: Do we need to use "innerxml", or can we actually use "chardata"?
	Value string `xml:",chardata"`
}

func (p PolicySetIdReference) Validate(errs *Errors) {
	if p.Value == "" {
		errs.Addf("Value not given")
	}
}
