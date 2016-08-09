package xacml

import "encoding/xml"

// 5.31: The <AttributeValue> element SHALL contain a literal attribute value.
type AttributeValue struct {
	Expression
	XMLName xml.Name `xml:"AttributeValue"`

	// The data-type of the attribute value.
	DataType string `xml:",attr"`
}

func (a AttributeValue) Validate(errs *Errors) {
	if a.DataType == "" {
		errs.Addf("DataType not given")
	}

	a.Expression.Validate(errs)
}
