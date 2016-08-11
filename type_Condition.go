package xacml

import "encoding/xml"

// 5.26: The <Condition> element is a Boolean function over attributes or
// functions of attributes.
type Condition struct {
	XMLName xml.Name `xml:"Condition"`

	// The <Condition> contains one <Expression> element, with the
	// restriction that the <Expression> return data-type MUST be
	// “http://www.w3.org/2001/XMLSchema#boolean”.  Evaluation of the
	// <Condition> element is described in Section 7.9.
	Expression ExpressionType
}

func (c *Condition) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	exp, err := ParseExpressions(decoder, start, nil)
	if err != nil {
		return err
	}

	if len(exp) > 0 {
		c.Expression = exp[0]
	}

	return nil
}

func (c Condition) Validate(errs *Errors) {
	if c.Expression == nil {
		errs.Addf("Expression not given")
	} else {
		val := c.Expression.(Validatable)
		val.Validate(errs.Sub("Expression"))
	}
}
