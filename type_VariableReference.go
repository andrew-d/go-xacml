package xacml

import "encoding/xml"

// 5.24: The <VariableReference> element is used to reference a value defined
// within the same encompassing <Policy> element.  The <VariableReference>
// element SHALL refer to the <VariableDefinition> element by identifier
// equality on the value of their respective VariableId attributes.  One and
// only one <VariableDefinition> MUST exist within the same encompassing
// <Policy> element to which the <VariableReference> refers.  There MAY be zero
// or more <VariableReference> elements that refer to the same
// <VariableDefinition> element.
type VariableReference struct {
	XMLName xml.Name `xml:"VariableReference"`

	// The name used to refer to the value defined in a
	// <VariableDefinition> element.
	VariableId string `xml:",attr"`
}

func (v VariableReference) Validate(errs *Errors) {
	if v.VariableId == "" {
		errs.Addf("VariableId not given")
	}
}
