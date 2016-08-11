package xacml

import "encoding/xml"

// 5.50: The <MultiRequests> element contains a list of requests by reference
// to <Attributes> elements in the enclosing <Request> element. The semantics
// of this element are defined in [Multi]. Support for this element is
// optional. If an implementation does not support this element, but receives
// it, the implementation MUST generate an “Indeterminate” response.
type MultiRequests struct {
	XMLName xml.Name `xml:"MultiRequests"`

	// Defines a request instance by reference to <Attributes> elements in
	// the enclosing <Request> element. See section 5.51.
	RequestReferences []RequestReference `xml:"RequestReference"`
}

func (m MultiRequests) Validate(errs *Errors) {
	if len(m.RequestReferences) < 1 {
		errs.Addf("RequestReferences not given")
	} else {
		for i, el := range m.RequestReferences {
			el.Validate(errs.SubN("RequestReferences", i))
		}
	}
}
