package xacml

import "encoding/xml"

// 5.8: The <AllOf> element SHALL contain a conjunctive sequence of <Match>
// elements.
type AllOf struct {
	XMLName xml.Name `xml:"AllOf"`

	// A conjunctive sequence of individual matches of the attributes in
	// the request context and the embedded attribute values.  See Section
	// 5.9.
	Matches []Match `xml:"Match"`
}

func (a AllOf) Validate(errs *Errors) {
	if len(a.Matches) < 1 {
		errs.Addf("Matches not given")
		return
	}

	for i, match := range a.Matches {
		match.Validate(errs.SubN("Matches", i))
	}
}
