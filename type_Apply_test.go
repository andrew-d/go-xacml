package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Apply(t *testing.T) {
	input := `
	<Apply FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-is-in">
		<Description>description here</Description>

		<AttributeValue
			DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>

		<AttributeDesignator
			AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:some-attribute"
			Category="urn:oasis:names:tc:xacml:1.0:subject-category:access-subject"
			DataType="http://www.w3.org/2001/XMLSchema#string"
			MustBePresent="true"/>
	</Apply>`

	dest := &Apply{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Apply")

	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:function:string-is-in", dest.FunctionId)
	assert.Equal(t, "description here", dest.Description)

	if assert.Len(t, dest.Arguments, 2) {
		assert.Equal(t, &AttributeValue{
			XMLName:  xml.Name{Local: "AttributeValue"},
			DataType: "http://www.w3.org/2001/XMLSchema#string",
			Value:    "value1",
		}, dest.Arguments[0])

		mbp := true
		assert.Equal(t, &AttributeDesignator{
			XMLName:       xml.Name{Local: "AttributeDesignator"},
			AttributeId:   "urn:oasis:names:tc:xacml:2.0:conformance-test:some-attribute",
			Category:      "urn:oasis:names:tc:xacml:1.0:subject-category:access-subject",
			DataType:      "http://www.w3.org/2001/XMLSchema#string",
			MustBePresent: &mbp,
		}, dest.Arguments[1])
	}
}
