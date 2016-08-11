package xacml

import "encoding/xml"

// 5.52: The <AttributesReference> element makes a reference to an <Attributes>
// element. The meaning of this element is defined in [Multi]. Support for this
// element is optional.
type AttributesReference struct {
	XMLName xml.Name `xml:"AttributesReference"`

	// A reference to the xml:id attribute of an <Attributes> element in
	// the enclosing <Request> element.
	ReferenceId string `xml:",attr"`
}

func (a AttributesReference) Validate(errs *Errors) {
	if a.ReferenceId == "" {
		errs.Addf("ReferenceId not given")
	}
}
