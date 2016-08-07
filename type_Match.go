package xacml

import "encoding/xml"

// 5.9: The <Match> element SHALL identify a set of entities by matching
// attribute values in an <Attributes> element of the request context with the
// embedded attribute value.
type Match struct {
	XMLName xml.Name `xml:"Match"`

	// Specifies a matching function.  The value of this attribute MUST be of
	// type xs:anyURI with legal values documented in Section 7.6.
	MatchId string `xml:",attr"` // TODO: required

	// Embedded attribute value.
	AttributeValue AttributeValue // TODO: required

	// TODO: one of the following two is required

	// MAY be used to identify one or more attribute values in an <Attributes>
	// element of the request context.
	AttributeDesignator *AttributeDesignator

	// MAY be used to identify one or more attribute values in a <Content>
	// element of the request context.
	AttributeSelector *AttributeSelector
}
