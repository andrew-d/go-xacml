package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIssuer(t *testing.T) {
	input := `
	<PolicyIssuer>
		<Content>foo</Content>
		<Attribute
			IncludeInResult="true"
			AttributeId="urn:oasis:names:tc:xacml:1.0:subject:subject-id">

			<AttributeValue
				DataType="http://www.w3.org/2001/XMLSchema#string">val1</AttributeValue>
		</Attribute>
	</PolicyIssuer>
	`

	dest := &PolicyIssuer{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIssuer")

	assert.Equal(t, "foo", dest.Content.XML)
	if assert.Len(t, dest.Attributes, 1) {
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:subject:subject-id", dest.Attributes[0].AttributeId)
	}
}
