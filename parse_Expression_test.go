package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expressionContainer struct {
	XMLName     xml.Name `xml:"Container"`
	Expressions []ExpressionType
	Err         error
}

func (c *expressionContainer) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	c.Expressions, c.Err = ParseExpressions(decoder, start, nil)
	return nil
}

func TestParseExpressions(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected interface{}
	}{
		{
			Input: `<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">doctor</AttributeValue>`,
			Expected: &AttributeValue{
				XMLName:  xml.Name{Local: "AttributeValue"},
				DataType: "http://www.w3.org/2001/XMLSchema#string",
				Value:    "doctor",
			},
		},
		{
			Input: `
			<AttributeDesignator
				AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
				DataType="http://www.w3.org/2001/XMLSchema#string"
				MustBePresent="false"
				/>
			`,
			Expected: &AttributeDesignator{
				XMLName:       xml.Name{Local: "AttributeDesignator"},
				AttributeId:   "urn:oasis:names:tc:xacml:1.0:action:action-id",
				DataType:      "http://www.w3.org/2001/XMLSchema#string",
				MustBePresent: new(bool),
			},
		},
		{
			Input: `<VariableReference VariableId="isFooBar"/>`,
			Expected: &VariableReference{
				XMLName:    xml.Name{Local: "VariableReference"},
				VariableId: "isFooBar",
			},
		},
	}

	for _, tc := range testCases {
		s := "<Container>" + tc.Input + "</Container>"
		dest := &expressionContainer{}

		err := xml.Unmarshal([]byte(s), dest)
		assert.NoError(t, err, "Error unmarshalling input for: %s", tc.Input)
		assert.NoError(t, dest.Err, "Error parsing Expressions for: %s", tc.Input)
		assert.Equal(t, tc.Expected, dest.Expressions[0])
	}
}
