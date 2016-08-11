package xacml

import "encoding/xml"

// 5.51: The <RequestReference> element defines an instance of a request in
// terms of references to <Attributes> elements. The semantics of this element
// are defined in [Multi]. Support for this element is optional.
type RequestReference struct {
	XMLName xml.Name `xml:"RequestReference"`

	// A reference to an <Attributes> element in the enclosing <Request>
	// element. See section 5.52.
	AttributesReferences []AttributesReference `xml:"AttributesReference"`
}

func (r RequestReference) Validate(errs *Errors) {
	for i, el := range r.AttributesReferences {
		el.Validate(errs.SubN("AttributesReference", i))
	}
}
