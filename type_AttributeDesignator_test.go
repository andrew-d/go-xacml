package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeDesignator(t *testing.T) {
	input := `
	<AttributeDesignator
		AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
		Category="urn:oasis:names:tc:xacml:3.0:attribute-category:action"
		DataType="http://www.w3.org/2001/XMLSchema#string"
		Issuer="an-issuer"
		MustBePresent="false"
		/>`

	dest := &AttributeDesignator{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeDesignator")

	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:action:action-id", dest.AttributeId)
	assert.Equal(t, "urn:oasis:names:tc:xacml:3.0:attribute-category:action", dest.Category)
	assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", dest.DataType)
	assert.Equal(t, "an-issuer", *dest.Issuer)
	assert.Equal(t, false, *dest.MustBePresent)
}
