package xacml

import "encoding/xml"

// 5.25: The <Expression> element is not used directly in a policy.  The
// <Expression> element signifies that an element that extends the
// ExpressionType and is a member of the <Expression> element substitution
// group SHALL appear in its place.
type Expression struct {
	XMLName xml.Name `xml:"Expression"`

	// The value
	// TODO: Do we need to use "innerxml", or can we actually use "chardata"?
	Value string `xml:",chardata"`
}

func (e Expression) Validate(errs *Errors) {
}
