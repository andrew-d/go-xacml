package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type AdviceExpressions struct {
	XMLName xml.Name `xml:"AdviceExpressions"`

	// Insert fields here
}

func (a AdviceExpressions) Validate(errs *Errors) {
}
