package xacml

import "encoding/xml"

// 5.17: The <CombinerParameter> element conveys a single parameter for a
// policy- or rule-combining algorithm.  Support for the <CombinerParameter>
// element is optional.
type CombinerParameter struct {
	XMLName xml.Name `xml:"CombinerParameter"`

	// The identifier of the parameter.
	ParameterName string `xml:",attr"`

	// The value of the parameter.
	AttributeValues []AttributeValue `xml:"AttributeValue"`
}

func (c CombinerParameter) Validate(errs *Errors) {
	if c.ParameterName == "" {
		errs.Addf("ParameterName not given")
	}

	if len(c.AttributeValues) < 1 {
		errs.Addf("AttributeValues not given")
	} else {
		for i, el := range c.AttributeValues {
			el.Validate(errs.SubN("AttributeValues", i))
		}
	}
}
