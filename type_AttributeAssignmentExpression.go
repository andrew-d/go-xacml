package xacml

import "encoding/xml"

// 5.41: The <AttributeAssignmentExpression> element is used for including
// arguments in obligations and advice.  It SHALL contain an AttributeId and an
// expression which SHALL by evaluated into the corresponding attribute value.
// The value specified SHALL be understood by the PEP, but it is not further
// specified by XACML. See Section 7.18.  Section 4.2.4.3 provides a number of
// examples of arguments included in obligations.
type AttributeAssignmentExpression struct {
	XMLName xml.Name `xml:"AttributeAssignmentExpression"`

	// The attribute identifier. The value of the AttributeId attribute in
	// the resulting <AttributeAssignment> element MUST be equal to this
	// value.
	AttributeId string `xml:",attr"`

	// An optional category of the attribute. If this attribute is missing,
	// the attribute has no category. The value of the Category attribute
	// in the resulting <AttributeAssignment> element MUST be equal to this
	// value.
	Category *string `xml:",attr"`

	// An optional issuer of the attribute. If this attribute is missing,
	// the attribute has no issuer. The value of the Issuer attribute in
	// the resulting <AttributeAssignment> element MUST be equal to this
	// value.
	Issuer *string `xml:",attr"`

	// The expression which evaluates to a constant attribute value or a
	// bag of zero or more attribute values. See section 5.25.
	Expression ExpressionType
}

func (a *AttributeAssignmentExpression) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "AttributeId":
			a.AttributeId = attr.Value
		case "Category":
			a.Category = &attr.Value
		case "Issuer":
			a.Issuer = &attr.Value
		}
	}

	exp, err := ParseExpressions(decoder, start)
	if err != nil {
		return err
	}

	if len(exp) > 0 {
		a.Expression = exp[0]
	}

	return nil
}

func (a AttributeAssignmentExpression) Validate(errs *Errors) {
	if a.AttributeId == "" {
		errs.Addf("AttributeId not given")
	}
	if a.Category != nil && *a.Category == "" {
		errs.Addf("Category given, but cannot be empty")
	}
	if a.Issuer != nil && *a.Issuer == "" {
		errs.Addf("Issuer cannot be empty")
	}
	if a.Expression == nil {
		errs.Addf("Expression not given")
	}
}
