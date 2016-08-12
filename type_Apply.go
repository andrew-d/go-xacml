package xacml

import (
	"encoding/xml"
	"fmt"
)

// 5.27: The <Apply> element denotes application of a function to its
// arguments, thus encoding a function call.  The <Apply> element can be
// applied to any combination of the members of the <Expression> element
// substitution group.  See Section 5.25.
type Apply struct {
	XMLName xml.Name `xml:"Apply"`

	// The identifier of the function to be applied to the arguments.
	// XACML-defined functions are described in Appendix A.3.
	FunctionId string `xml:",attr"`

	// A free-form description of the <Apply> element.
	Description string `xml:"Description,omitempty"`

	// Arguments to the function, which may include other functions.
	Arguments []ExpressionType `xml:",any"`
}

func (a *Apply) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "FunctionId":
			a.FunctionId = attr.Value
		}
	}

	exp, err := ParseExpressions(decoder, start, func(tok xml.StartElement) error {
		if tok.Name.Local == "Description" {
			type DummyDescription struct {
				XMLName xml.Name `xml:"Description"`
				Value   string   `xml:",chardata"`
			}

			desc := &DummyDescription{}
			if err := decoder.DecodeElement(desc, &tok); err != nil {
				return err
			}

			a.Description = desc.Value
			return nil
		}

		return fmt.Errorf("unknown element does not implement <Expression>: %s", tok.Name.Local)
	})
	if err != nil {
		return err
	}

	a.Arguments = exp
	return nil
}

func (a Apply) Validate(errs *Errors) {
	if a.FunctionId == "" {
		errs.Addf("FunctionId not given")
	}
	if len(a.Arguments) == 0 {
		errs.Addf("No arguments given")
	}

	for i, el := range a.Arguments {
		val := el.(Validatable)
		val.Validate(errs.SubN("Arguments", i))
	}
}
