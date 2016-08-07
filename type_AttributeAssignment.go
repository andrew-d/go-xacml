package xacml

import "encoding/xml"

// 5.36: The <AttributeAssignment> element is used for including arguments in
// obligation and advice expressions.  It SHALL contain an AttributeId and the
// corresponding attribute value, by extending the AttributeValueType type
// definition.  The <AttributeAssignment> element MAY be used in any way that
// is consistent with the schema syntax, which is a sequence of <xs:any>
// elements. The value specified SHALL be understood by the PEP, but it is not
// further specified by XACML. See Section 7.18.  Section 4.2.4.3 provides a
// number of examples of arguments included in obligation.expressions.
type AttributeAssignment struct {
	AttributeValue
	XMLName xml.Name `xml:"AttributeAssignment"`

	// The attribute identifier
	AttributeId string `xml:",attr"`

	// An optional category of the attribute. If this attribute is missing, the
	// attribute has no category. The PEP SHALL interpret the significance and
	// meaning of any Category attribute. Non-normative note: an expected use
	// of the category is to disambiguate attributes which are relayed from the
	// request.
	Category *string `xml:",attr"`

	// An optional issuer of the attribute. If this attribute is missing, the
	// attribute has no issuer. The PEP SHALL interpret the significance and
	// meaning of any Issuer attribute. Non-normative note: an expected use of
	// the issuer is to disambiguate attributes which are relayed from the
	// request.
	Issuer *string `xml:",attr"`
}

func (a AttributeAssignment) Validate(errs *Errors) {
	if a.AttributeId == "" {
		errs.Addf("AttributeId not given")
	}
	if a.Category != nil && *a.Category == "" {
		errs.Addf("Category given, but cannot be empty")
	}
	if a.Issuer != nil && *a.Issuer == "" {
		errs.Addf("Issuer given, but cannot be empty")
	}

	// Note: not usnig errs.Sub, since this is embedded
	a.AttributeValue.Validate(errs)
}
