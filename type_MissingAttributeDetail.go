package xacml

import "encoding/xml"

// 5.58: The <MissingAttributeDetail> element conveys information about
// attributes required for policy evaluation that were missing from the request
// context.
type MissingAttributeDetail struct {
	XMLName xml.Name `xml:"MissingAttributeDetail"`

	// The category identifier of the missing attribute.
	Category string `xml:",attr"`

	// The identifier of the missing attribute.
	AttributeId string `xml:",attr"`

	// The data-type of the missing attribute.
	DataType string `xml:",attr"`

	// This attribute, if supplied, SHALL specify the required Issuer of
	// the missing attribute.
	Issuer *string `xml:",attr"`

	// The required value of the missing attribute.
	AttributeValue AttributeValue
}

func (m MissingAttributeDetail) Validate(errs *Errors) {
	if m.Category == "" {
		errs.Addf("Category not given")
	}
	if m.AttributeId == "" {
		errs.Addf("AttributeId not given")
	}
	if m.DataType == "" {
		errs.Addf("DataType not given")
	}
	if m.Issuer != nil && *m.Issuer == "" {
		errs.Addf("Issuer given, but cannot be empty")
	}

	m.AttributeValue.Validate(errs.Sub("AttributeValue"))
}
