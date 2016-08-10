package xacml

import (
	"encoding/xml"
	"fmt"
)

// An alias for an XML element / struct that is a valid Expression.  From
// section 7.4, the appropriate elements are:
//   - <xacml:AttributeValue>
//   - <xacml:AttributeDesignator>
//   - <xacml:AttributeSelector>
//   - <xacml:Apply>
//   - <xacml:Function>
//   - <xacml:VariableReference>
//
// This package will never create an ExpressionType that is not one of the
// corresponding elements.
type ExpressionType interface{}

// Starting from the given element, will attempt to parse children that
// implement <Expression>.  Note that this does not parse attributes on the
// given start element.
func ParseExpressions(decoder *xml.Decoder, start xml.StartElement, unknown func(xml.StartElement) error) ([]ExpressionType, error) {
	if unknown == nil {
		unknown = func(tok xml.StartElement) error {
			return fmt.Errorf("unknown element does not implement <Expression>: %s", tok.Name.Local)
		}
	}

	ret := []ExpressionType{}

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch tok := token.(type) {
		case xml.StartElement:
			var elementStruct interface{}

			switch tok.Name.Local {
			case "AttributeValue":
				elementStruct = &AttributeValue{}
			case "AttributeDesignator":
				elementStruct = &AttributeDesignator{}
			case "AttributeSelector":
				elementStruct = &AttributeSelector{}
			case "Apply":
				elementStruct = &Apply{}
			case "Function":
				elementStruct = &Function{}
			case "VariableReference":
				elementStruct = &VariableReference{}
			default:
				if err := unknown(tok); err != nil {
					return nil, err
				}
				continue
			}

			if err = decoder.DecodeElement(elementStruct, &tok); err != nil {
				return nil, err
			}

			ret = append(ret, elementStruct)

		case xml.EndElement:
			return ret, nil
		}
	}
}
