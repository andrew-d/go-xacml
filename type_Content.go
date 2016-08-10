package xacml

import "encoding/xml"

// 5.45: The <Content> element is a notional placeholder for additional
// attributes, typically the content of the resource.
type Content struct {
	XMLName xml.Name `xml:"Content"`

	// The <Content> element has exactly one arbitrary type child element.
	XML string `xml:",innerxml"`
}

func (c Content) Validate(errs *Errors) {
	// TODO: Validate single child
}
