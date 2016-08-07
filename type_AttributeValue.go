package xacml

import "encoding/xml"

// 5.31: The <AttributeValue> element SHALL contain a literal attribute value.
type AttributeValue struct {
	XMLName xml.Name `xml:"AttributeValue"`

	// The data-type of the attribute value.
	DataType string `xml:",attr"`

	// The value
	// TODO: Do we need to use "innerxml", or can we actually use "chardata"?
	Value string `xml:",chardata"`
}

func (a AttributeValue) Validate(errs *Errors) {
	if a.DataType == "" {
		errs.Addf("DataType not given")
	}
}
