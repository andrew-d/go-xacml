package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Attributes(t *testing.T) {
	input := `
	<Attributes
		xml:id="subject1"
		Category="urn:oasis:names:tc:xacml:1.0:subject-category:access-subject">

		<Attribute
			IncludeInResult="true"
			AttributeId="urn:oasis:names:tc:xacml:1.0:subject:subject-id">

			<AttributeValue
				DataType="http://www.w3.org/2001/XMLSchema#string">val1</AttributeValue>
		</Attribute>

		<Content>foo bar</Content>
	</Attributes>`

	dest := &Attributes{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Attributes")

	assert.Equal(t, "subject1", dest.Id)
	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:subject-category:access-subject", dest.Category)

	if assert.Len(t, dest.Attributes, 1) {
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:subject:subject-id", dest.Attributes[0].AttributeId)
	}

	assert.Equal(t, "foo bar", dest.Content.XML)
}
