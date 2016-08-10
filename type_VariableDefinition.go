package xacml

import "encoding/xml"

// 5.23: The <VariableDefinition> element SHALL be used to define a value that
// can be referenced by a <VariableReference> element.  The name supplied for
// its VariableId attribute SHALL NOT occur in the VariableId attribute of any
// other <VariableDefinition> element within the encompassing policy.  The
// <VariableDefinition> element MAY contain undefined <VariableReference>
// elements, but if it does, a corresponding <VariableDefinition> element MUST
// be defined later in the encompassing policy.  <VariableDefinition> elements
// MAY be grouped together or MAY be placed close to the reference in the
// encompassing policy.  There MAY be zero or more references to each
// <VariableDefinition> element.
type VariableDefinition struct {
	XMLName xml.Name `xml:"VariableDefinition"`

	// The name of the variable definition.
	VariableId string `xml:",attr"`

	// Any element of ExpressionType complex type.
	Expression ExpressionType
}

func (v *VariableDefinition) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "VariableId":
			v.VariableId = attr.Value
		}
	}

	exp, err := ParseExpressions(decoder, start, nil)
	if err != nil {
		return err
	}

	if len(exp) > 0 {
		v.Expression = exp[0]
	}

	return nil
}

func (v VariableDefinition) Validate(errs *Errors) {
	if v.VariableId == "" {
		errs.Addf("VariableId not given")
	}

	if v.Expression == nil {
		errs.Addf("Expression not given")
	} else {
		val := v.Expression.(Validatable)
		val.Validate(errs.Sub("Expression"))
	}
}
