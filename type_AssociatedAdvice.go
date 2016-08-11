package xacml

import "encoding/xml"

// 5.33: The <AssociatedAdvice> element SHALL contain a set of <Advice>
// elements.
type AssociatedAdvice struct {
	XMLName xml.Name `xml:"AssociatedAdvice"`

	// A sequence of advice.  See Section 5.35.
	Advice []Advice `xml:"Advice"`
}

func (a AssociatedAdvice) Validate(errs *Errors) {
	if len(a.Advice) < 1 {
		errs.Addf("Advice not given")
	} else {
		for i, el := range a.Advice {
			el.Validate(errs.SubN("Advice", i))
		}
	}
}
