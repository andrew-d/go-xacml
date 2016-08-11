package xacml

import "encoding/xml"

// 5.32: The <Obligations> element SHALL contain a set of <Obligation>
// elements.
type Obligations struct {
	XMLName xml.Name `xml:"Obligations"`

	// A sequence of obligations.  See Section 5.34.
	Obligations []Obligation `xml:"Obligation"`
}

func (o Obligations) Validate(errs *Errors) {
	if len(o.Obligations) < 1 {
		errs.Addf("Obligations not given")
	} else {
		for i, el := range o.Obligations {
			el.Validate(errs.SubN("Obligations", i))
		}
	}
}
