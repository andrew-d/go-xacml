package xacml

import "encoding/xml"

// 5.48: The <Result> element represents an authorization decision result.  It
// MAY include a set of obligations that MUST be fulfilled by the PEP.  If the
// PEP does not understand or cannot fulfill an obligation, then the action of
// the PEP is determined by its bias, see section 7.1. It MAY include a set of
// advice with supplemental information which MAY be safely ignored by the PEP.
type Result struct {
	XMLName xml.Name `xml:"Result"`

	// TODO: finish me
}

func (r Result) Validate(errs *Errors) {
	// TODO: finish me
}
