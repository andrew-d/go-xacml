package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Attribute(t *testing.T) {
	input := `
	<Attribute
		IncludeInResult="true"
		AttributeId="urn:oasis:names:tc:xacml:1.0:subject:subject-id">

		<AttributeValue
			DataType="http://www.w3.org/2001/XMLSchema#string">John Doe</AttributeValue>
	</Attribute>`

	dest := &Attribute{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Attribute")

	assert.Equal(t, true, dest.IncludeInResult)
	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:subject:subject-id", dest.AttributeId)

	if assert.Len(t, dest.AttributeValues, 1) {
		assert.Equal(t, "John Doe", dest.AttributeValues[0].Value)
	}
}

func TestAttributeNoValues(t *testing.T) {
	input := `<Attribute AttributeId="fake"></Attribute>`

	dest := &Attribute{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Attribute")

	errs := NewErrors("Attribute")
	dest.Validate(errs)
	assert.Equal(t, 1, errs.NumErrors())
	assert.Contains(t, errs.Summary(), "AttributeValues not given")
}
