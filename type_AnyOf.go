package xacml

import "encoding/xml"

// 5.7: The <AnyOf> element SHALL contain a disjunctive sequence of <AllOf>
// elements.
type AnyOf struct {
	XMLName xml.Name `xml:"AnyOf"`

	// See Section 5.8.
	AllOf []AllOf `xml:"AllOf"`
}

func (a AnyOf) Validate(errs *Errors) {
	if len(a.AllOf) < 1 {
		errs.Addf("AllOf not given")
		return
	}

	for i, all := range a.AllOf {
		all.Validate(errs.SubN("AllOf", i))
	}
}
