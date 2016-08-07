package xacml

import "encoding/xml"

// 5.29: The <AttributeDesignator> element retrieves a bag of values for a
// named attribute from the request context.  A named attribute SHALL be
// considered present if there is at least one attribute that matches the
// criteria set out below.
//
// The <AttributeDesignator> element SHALL return a bag containing all the
// attribute values that are matched by the named attribute.  In the event that
// no matching attribute is present in the context, the MustBePresent attribute
// governs whether this element returns an empty bag or “Indeterminate”.  See
// Section 7.3.5.
//
// The <AttributeDesignator> MAY appear in the <Match> element and MAY be
// passed to the <Apply> element as an argument.
//
// The <AttributeDesignator> element is of the AttributeDesignatorType complex
// type.
type AttributeDesignator struct {
	XMLName xml.Name `xml:"AttributeDesignator"`

	// This attribute SHALL specify the AttributeId with which to match the
	// attribute.
	AttributeId string `xml:",attr"` // TODO: required

	// This attribute SHALL specify the Category with which to match the
	// attribute.
	Category string `xml:",attr"` // TODO: required

	// The bag returned by the <AttributeDesignator> element SHALL contain
	// values of this data-type.
	DataType string `xml:",attr"` // TODO: required

	// This attribute, if supplied, SHALL specify the Issuer with which to
	// match the attribute.
	Issuer string `xml:",attr"`

	// This attribute governs whether the element returns “Indeterminate” or an
	// empty bag in the event the named attribute is absent from the request
	// context.  See Section 7.3.5.  Also see Sections 7.19.2 and 7.19.3.
	MustBePresent bool `xml:",attr"` // TODO: required
}
