package xacml

import "encoding/xml"

// 5.9: The <Match> element SHALL identify a set of entities by matching
// attribute values in an <Attributes> element of the request context with the
// embedded attribute value.
type Match struct {
	XMLName xml.Name `xml:"Match"`

	// Specifies a matching function.  The value of this attribute MUST be of
	// type xs:anyURI with legal values documented in Section 7.6.
	MatchId string `xml:",attr"`

	// Embedded attribute value.
	AttributeValue AttributeValue

	// MAY be used to identify one or more attribute values in an <Attributes>
	// element of the request context.
	AttributeDesignator *AttributeDesignator

	// MAY be used to identify one or more attribute values in a <Content>
	// element of the request context.
	AttributeSelector *AttributeSelector
}

func (m Match) Validate(errs *Errors) {
	if m.MatchId == "" {
		errs.Addf("MatchId not given")
	}

	m.AttributeValue.Validate(errs.Sub("AttributeValue"))

	// Only one of these two is allowed
	ad := m.AttributeDesignator != nil
	as := m.AttributeSelector != nil

	if ad && as {
		errs.Addf("AttributeDesignator and AttributeSelector cannot both be provided")
	} else if !ad && !as {
		errs.Addf("neither AttributeDesignator nor AttributeSelector given")
	} else if ad {
		m.AttributeDesignator.Validate(errs.Sub("AttributeDesignator"))
	} else {
		m.AttributeSelector.Validate(errs.Sub("AttributeSelector"))
	}
}
