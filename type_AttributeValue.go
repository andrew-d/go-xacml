package xacml

import "encoding/xml"

// 5.31: The <AttributeValue> element SHALL contain a literal attribute value.
type AttributeValue struct {
	XMLName xml.Name `xml:"AttributeValue"`

	// The data-type of the attribute value.
	DataType string `xml:",attr"`

	// From Appendix A.2: When the value is encoded in an <AttributeValue>
	// element, the namespace context is given by the <AttributeValue>
	// element and an XML attribute called XPathCategory gives the category
	// of the <Content> element where the expression applies.
	XPathCategory string `xml:",attr,omitempty"`

	// TODO: Do we need to use "innerxml", or can we actually use "chardata"?
	Value string `xml:",chardata"`
}

func (a AttributeValue) Validate(errs *Errors) {
	if a.DataType == "" {
		errs.Addf("DataType not given")
	}
}
