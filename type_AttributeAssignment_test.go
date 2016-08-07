package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeAssignment(t *testing.T) {
	input := `
	<AttributeAssignment
		AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1"
		DataType="http://www.w3.org/2001/XMLSchema#string"
		Category="urn:oasis:names:tc:xacml:3.0:attribute-category:resource"
		Issuer="an-issuer">assignment1</AttributeAssignment>
	`

	dest := &AttributeAssignment{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeAssignment")

	// Insert specific tests here
	assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1", dest.AttributeId)
	assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", dest.DataType)
	assert.Equal(t, "urn:oasis:names:tc:xacml:3.0:attribute-category:resource", *dest.Category)
	assert.Equal(t, "an-issuer", *dest.Issuer)

	assert.Equal(t, "assignment1", dest.Value)
}
